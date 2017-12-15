package main

import (
  "hash/fnv"
  "sort"
  "fmt"
)

type BigBucket struct {
  Name    string
  Buckets []*Bucket
}

type Bucket struct {
  Big      *BigBucket
  Position uint64
}

type ConsistentRing struct {
  Range   uint64
  Buckets []Bucket
  BigBuckets BigBucketsCollection
}

type BigBucketsCollection []*BigBucket

func (a ConsistentRing) Len() int      { return len(a.Buckets) }
func (a ConsistentRing) Swap(i, j int) { a.Buckets[i], a.Buckets[j] = a.Buckets[j], a.Buckets[i] }
func (a ConsistentRing) Less(i, j int) bool {
  return (a.Buckets[i].Position) < (a.Buckets[j].Position)
}
// sort.Sort c.BigBucket 便于search, 由于Bucket有BigBucket指针, 废弃
// func (a BigBucketsCollection) Len() int      { return len(a) }
// func (a BigBucketsCollection) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// func (a BigBucketsCollection) Less(i, j int) bool {
//   return (*a[i]).Name < (*a[j]).Name
// }

// 支持BigBucket
func (c *ConsistentRing) AddBigBucket(name string) BigBucket {
  h := fnv.New64()
  big_bucket := BigBucket{name, make([]*Bucket, 0, 3)}
  for i := 1; i <= 3; i++ {
    h.Write([]byte(name))
    bucket := Bucket{&big_bucket, uint64(h.Sum64()) % c.Range}
    c.Buckets = append(c.Buckets, bucket)
    big_bucket.Buckets = append(big_bucket.Buckets, &bucket)
  }
  // sort.Sort(BigBucketsCollection(c.BigBuckets))  // 由于Bucket有BigBucket指针, 废弃
  sort.Sort(c)
  return big_bucket
}

// 支持BigBucket
func (c ConsistentRing) DumpNodesRange() ConsistentRing {
  sort.Sort(c)
  fmt.Printf("%v\n", c)
  return c
}

func (c ConsistentRing) FindBigBucketByKey(key string) (b BigBucket) {
  bucket := c.FindBucketByKey(key)
  return *bucket.Big
}

func (c ConsistentRing) FindBucketByKey(key string) (b Bucket) {
  keyh := fnv.New64()
  keyh.Write([]byte(key))
  key_pos := keyh.Sum64() % c.Range
  start_bucket_idx := (sort.Search(len(c.Buckets), func(i int) bool {
    return c.Buckets[i].Position > key_pos
  }) + len(c.Buckets) - 1) % len(c.Buckets)
  b = c.Buckets[start_bucket_idx]
  // start := b.Position
  // end := c.Buckets[(start_bucket_idx+1)%len(c.Buckets)].Position
  // fmt.Printf("%d, start: %d, end: %d\n", key_pos, start, end)
  return
}

func main() {
  c := ConsistentRing{100, make([]Bucket, 0, 10), make(BigBucketsCollection, 10)}
  b := c.AddBigBucket("bbb")
  c.AddBigBucket("ccc")
  c.AddBigBucket("ddd")
  c.AddBigBucket("eee")
  a := c.AddBigBucket("aaa")

  result_aaa := c.FindBigBucketByKey("aaa")
  result_bbb := c.FindBigBucketByKey("bbb")

  fmt.Printf("BigBucket %v\n", a)
  fmt.Printf("%v\n", result_aaa)
  fmt.Printf("BigBucket %v\n", b)
  fmt.Printf("%v\n", result_bbb)
}
