package main

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// if
	pressKeys(kb, false, false, keybd_event.VK_I, keybd_event.VK_F)

	// space
	pressKeys(kb, false, false, keybd_event.VK_SPACE)

	// err
	pressKeys(kb, false, false, keybd_event.VK_E, keybd_event.VK_R, keybd_event.VK_R)

	// space
	pressKeys(kb, false, false, keybd_event.VK_SPACE)

	// !
	pressKeys(kb, true, false, keybd_event.VK_1)

	// =
	pressKeys(kb, false, false, keybd_event.VK_EQUAL)

	// space
	pressKeys(kb, false, false, keybd_event.VK_SPACE)

	// nil
	pressKeys(kb, false, false, keybd_event.VK_N, keybd_event.VK_I, keybd_event.VK_L)

	// space
	pressKeys(kb, false, false, keybd_event.VK_SPACE)

	// {
	pressKeys(kb, true, false, keybd_event.VK_SP4)

	// enter
	pressKeys(kb, false, false, keybd_event.VK_ENTER)

	// tab
	// pressKeys(kb, false, false, keybd_event.VK_TAB)

	// log
	pressKeys(kb, false, false, keybd_event.VK_L, keybd_event.VK_O, keybd_event.VK_G)

	// .
	pressKeys(kb, false, false, keybd_event.VK_DOT)

	// P
	pressKeys(kb, true, false, keybd_event.VK_P)

	// rintln
	pressKeys(kb, false, false, keybd_event.VK_R, keybd_event.VK_I, keybd_event.VK_N, keybd_event.VK_T, keybd_event.VK_L, keybd_event.VK_N)

	// (
	pressKeys(kb, true, false, keybd_event.VK_9)

	// err
	pressKeys(kb, false, false, keybd_event.VK_E, keybd_event.VK_R, keybd_event.VK_R)

	// .
	pressKeys(kb, false, false, keybd_event.VK_DOT)

	// E
	pressKeys(kb, true, false, keybd_event.VK_E)

	// rror
	pressKeys(kb, false, false, keybd_event.VK_R, keybd_event.VK_R, keybd_event.VK_O, keybd_event.VK_R)

	// (
	pressKeys(kb, true, false, keybd_event.VK_9)

	// )
	pressKeys(kb, true, false, keybd_event.VK_0)

	// )
	pressKeys(kb, true, false, keybd_event.VK_0)

	// enter
	// pressKeys(kb, false, false, keybd_event.VK_ENTER)

	// }
	// pressKeys(kb, true, false, keybd_event.VK_SP5)
}

func pressKeys(kb keybd_event.KeyBonding, shift bool, ctrl bool, keys ...int) {
	kb.HasSHIFT(shift)

	for _, key := range keys {
		kb.SetKeys(
			key,
		)

		// Or you can use Press and Release
		err := kb.Press()
		time.Sleep(10 * time.Millisecond)
		err = kb.Release()

		// Press the selected keys
		// err := kb.Launching()
		if err != nil {
			panic(err)
		}
	}
}
