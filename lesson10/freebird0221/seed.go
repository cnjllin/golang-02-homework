func (s *Seed) MakeSeed(rootPath string) (err error) {
	var totalFileSize int64
	var block_rem_size int64
	var file_idx int64

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		log.Debug(info.Mode()&os.ModeType == 0, path)
		if info.Mode()&os.ModeType != 0 {
			return nil
		}
		totalFileSize += info.Size()
		file_info := FileInfo{
			Path: path,
			Size: info.Size(),
			Mode: info.Mode(),
		}
		s.FileList = append(s.FileList, file_info)

		file_idx += 1
		if block_rem_size < file_info.Size {
			for block_rem_size < file_info.Size {
				s.BlockList = append(s.BlockList, BlockInfo{int64(file_idx), block_rem_size })
				s.BlockCount += 1
				block_rem_size += s.BlockSize
			}
		}
		block_rem_size -= file_info.Size

		return nil
	})

	//log.Debug(s.FileList)
	log.Debug("Print Seed ", s)
	log.Debug("totalFileSize ", totalFileSize)

	//var (
	//	totalBlockSize int64
	//	lastBlockIdx   int64 = -1
	//)
	/*
		+-----------------+-----+-----+-----+-----+-----+-----+-----+-
		|block|block|     |     |     |     |     |     |     |     |
		|  0  |  1  |  2  |  3  |  4  | ... |     |     |     |     |
		+-----+---+-+-----+-----+-----+-----+-----+-----+----++-----+-
		|  file0  | 1|2|3|file4|          Big file5          |  ...
		|         |  | | |     |                             |
		+---------+--+-+-+-----+-----------------------------+--------
	*/
	//for blockIdx, block := range s.BlockList {
	//	if totalFileSize > totalBlockSize {
	//		totalBlockSize += s.BlockSize
	//		for fileIdx, file := range s.FileList {
	//
	//		}
	//	}
	//}

	//s.BlockList = append(s.BlockList, BlockInfo{})

	return
}
