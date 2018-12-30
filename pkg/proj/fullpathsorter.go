package proj

type FullPathSorter []Project

func (p FullPathSorter) Len() int           { return len(p) }
func (p FullPathSorter) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p FullPathSorter) Less(i, j int) bool { return p[i].FullPath <= p[j].FullPath }
