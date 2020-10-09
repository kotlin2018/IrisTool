package utils

import "os"

//根据文件名创建文件，并将传入字符内容写入文件内
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
