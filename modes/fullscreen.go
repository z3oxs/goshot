
package modes

import "github.com/z3oxs/ghost/utils"

func Fullscreen(save, format string, clipboard, output, upload bool) {
    screenshot := utils.CaptureScreen(save)

    utils.PlaySound("screenshot.wav")
	
    if utils.CheckSave(save) {
       utils.SaveImage(screenshot, save) 

    }

    if clipboard {
        utils.SaveToClipboard(screenshot)
    
    }

    if output {
        utils.OutputToStdout(screenshot, format)

    }

    if upload {
        utils.UploadImage(screenshot, format)

    }
}

