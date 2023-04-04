/**
 * @Author: smono
 * @Description:
 * @File:  del
 * @Version: 1.0.0
 * @Date: 2022/9/28 22:53
 */

package fileBase

import "os"

func (p *FileBase) Del(filePath string) error {
	return os.Remove(filePath)
}
