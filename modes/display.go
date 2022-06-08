package modes

import (
    "log"

    "github.com/z3oxs/ghost/utils"
)

func Display(save, format string, clipboard, output, upload bool, display int) {
    displays := utils.GetDisplays()

    if display > len(displays) {
        log.Fatal("Invalid screen selected.")

    }

    displaySelected := displays[display]

    screenshot := utils.CaptureRect (
        displaySelected.X,
        displaySelected.Y,
        displaySelected.Width,
        displaySelected.Height,
        save,
    )

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
