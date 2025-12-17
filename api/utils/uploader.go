package utils

import (

	"os"
	"errors"
	"image"
	
	"path/filepath"
)
const uploadDir="uploads"
func  UplaodImage(filename string ,image image.Image)(error){
	err:=os.Mkdir(uploadDir,0755)
	if err!=nil{
		return  errors.New("error creating uploads directory")
	}
filepath:=filepath.Join(uploadDir,filename)
file,err:=os.Create(filepath)
if err!=nil{
	return  errors.New("error creating file")
}
defer file.Close()
	return  nil;
}
