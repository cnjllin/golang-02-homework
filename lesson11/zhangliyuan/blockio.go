package transfer

import (
	"bufio"
	log "github.com/auxten/logrus"
	"github.com/auxten/gink-go/src/seed"
	"os"
	"strings"
	"github.com/Sirupsen/logrus"
)

type BlockReader interface {
	// 1. 通过RPC告知对端要从blockIndex开始收blockCount个块
	// 2. 获取一个bufio.Reader
	Read(blockIndex int64, blockCount int64) (reader bufio.Reader, err error)
}

type BlockWriter interface {
	// 1. 通过blockindex确定应该在哪个文件的位置开始写
	Write(blockIndex int64, blockCount int64) (writer bufio.Writer, err error)
}

type BlockFileIO interface {
	BlockReader
	BlockWriter
}

type BlockSocketIO interface {
	BlockReader
	BlockWriter
}

type BlockServer struct {
	seed seed.Seed
	BlockFileIO
	BlockSocketIO
}

func (b *BlockServer) Read(blockIndex int64, blockCount int64) (reader bufio.Reader, err error) {

	return
}

func (b *BlockServer) Write(blockIndex int64, blockCount int64) (reader bufio.Writer, err error) {
	return
}

/*
	scp -r ./src_dir/src xx.com:./dst_dir/dst
	if src is Dir and dst is dir
		./dst_dir/dst/src
	if src is Dir and dst not exist
		./dst_dir/dst
	if src is Dir and dst is file
		fail
	if src is file and dst is file
		overwrite
	if src is file and dst is dir
		./dst_dir/dst/src

	scp -r ./src_dir/src xx.com:./dst_dir/dst/
	if src is Dir and dst is dir
		./dst_dir/dst/src
	if src is Dir and dst is file
		fail
	if src is file and dst is file
		fail
	if src is file and dst is dir
		./dst_dir/dst/src
 */

func (b *BlockServer) DownloadBlock(startBlockIndex int64, blockCount int64) (count int64, err error) {
	socketReader, err := b.BlockSocketIO.Read(startBlockIndex, blockCount)
	if err != nil {
		log.Errorf("get block from socket error, idx:%d, count:%d", startBlockIndex, blockCount)
	}

	var Previous_fileindex int64
	Previous_fileindex=0
	wblock := b.seed.BlockList[startBlockIndex]
	fileIndex := wblock.StartFileIndex
	fileOffset:=wblock.StartOffset
	wfile := b.seed.FileList[fileIndex]
	LABEL:
	for remainSize := b.seed.BlockSize * blockCount; remainSize > 0; {
		// todo 创建本文件所有需要的文件夹，defer-close

		filepath:=strings.Split(wfile.Path,string(os.PathSeparator))
		if fileOffset==0{
			x,err:=os.Create(filepath[len(filepath)-1])
			check(err)
			x.Close()
		}
		filename:=filepath[len(filepath)-1]
		f,err:=os.OpenFile(filename,os.O_CREATE|os.O_WRONLY,0644)
		check(err)
		defer f.Close()
		var less_size int64 //该文件剩余大小

		for index:=0;;index++{
			if b.seed.BlockList[index].StartFileIndex==fileIndex{
				less_size=b.seed.FileList[fileIndex].Size-fileOffset   //文件剩余大小=文件大小-偏移量
				if less_size<b.seed.BlockSize{      //如果 文件剩余大小<块大小,则写入文件大小-偏移量大小数据到文件
					f.Seek(fileOffset,0)
					f.Write([]byte("剩下的部分大小"))//filesize - offset
					f.Close()
					remainSize=remainSize-b.seed.FileList[fileIndex].Size   //循环参数remainsize=总长-文件长度

					Previous_fileindex=fileIndex		//标记该文件index,为下一个同块的文件index做比较判断用
					wfile=b.seed.FileList[fileIndex+1]	//该下个文件作为wfile了
					fileIndex+=1						//文件索引+1
					continue LABEL						//回到remain循环
				}else if less_size==b.seed.BlockSize{		//如果文件剩余大小 恰好等于 块大小
					f.Seek(fileOffset,0)				//写入该块大小的内容
					f.Write([]byte("剩下的部分大小"))//filesize - offset
					f.Close()
					remainSize=remainSize-b.seed.FileList[fileIndex].Size		//总长度-文件长度

					Previous_fileindex=0				//下一个块不需要同块文件比较
					wfile=b.seed.FileList[fileIndex+1]	//下个文件作为wfile
					fileIndex+=1						//文件索引+1
					wblock=b.seed.BlockList[index+1]	//wblock为下一块
					continue LABEL
				}
				f.Seek(fileOffset,0)				//如果文件剩余大小大于该块  该块全写进去 ,wblock为下一块
				f.Write([]byte("file"))   //往里写什么啊,文件内容来自块吗?
				wblock=b.seed.BlockList[index+1]
			}else if Previous_fileindex!=0 && fileIndex==b.seed.BlockList[index].StartFileIndex+1 {
				//如果 上一个previous不等于0 且 该块的startindex加1等于文件索引 代表上一个文件未用完全部块
				if wfile.Size<remainSize%b.seed.BlockSize{  //如果文件长度小于该块剩余的所有的部分,那么直接写入该文件长度内容,而且将fileindex
				//加1,即代表该块未结束 仍然有下个文件从此块开始
					f.Write([]byte(""))//写入file size 长度的文件内容
					f.Close()
					wfile=b.seed.FileList[fileIndex+1]  //设置下一个文件wfile
					Previous_fileindex=fileIndex  //fileindex不变,为了下个文件(第三个)继续在本块比对
					remainSize=remainSize-b.seed.FileList[fileIndex].Size
					continue LABEL

				}else if wfile.Size==remainSize%b.seed.BlockSize{//该文件大小与剩余块相同大小
					f.Write([]byte(""))//写入file size 也就是remainsize%b.seed.blocksize
					f.Close()
					remainSize=remainSize-b.seed.FileList[fileIndex].Size

					Previous_fileindex=0  //下一块就不是从半块开头了,所以重置回1
					wfile=b.seed.FileList[fileIndex+1] //文件下一个
					wblock=b.seed.BlockList[index+1] //块下一个
					fileIndex+=1 //索引下一个
					continue LABEL

				}else{
					f.Write([]byte("当前block剩余大小"))//remainsize%b.seed.blocksize
					wblock=b.seed.BlockList[index+1] //正常循环块
					Previous_fileindex=0 //正常块,设置为0
				}
			}else{
				continue
			}
		}
		/* todo
			打开文件，定位到块起始的位置，写入数据，不要超出原有文件大小
		 	写完一个文件就进行下一次循环，主要是fileIndex
			注意对remainSize进行减小
		*/


	}

	return
}

func check(err error){
	if err !=nil {
		logrus.Debug(err)
	}
}