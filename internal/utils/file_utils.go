package utils

import "os"

// Удалить файл
func DeleteFile(path string) error {
    return os.Remove(path)
    
}
