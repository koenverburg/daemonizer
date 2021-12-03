package utils

import (
	"fmt"
	"reflect"

	tmux "github.com/jubnzv/go-tmux"
)

func CreateTmuxSession() {
	server := new(tmux.Server)

	// Check that "example" session already exists.
	exists, err := server.HasSession("example")
	if err != nil {
		msg := fmt.Errorf("Can't check 'example' session: %s", err)
		fmt.Println(msg)
		return
	}

	if exists {
		// Sure, you can use KillSession here.
		fmt.Println("Session 'example' already exists!")
		fmt.Println("Please stop it before running this demo.")
		return
	}

	// Prepare configuration for a new session with some windows.
	session := tmux.Session{Name: "example-session"}

	w1 := tmux.Window{Name: "first", Id: 0}
	w2 := tmux.Window{Name: "second", Id: 1}

	session.AddWindow(w1)
	session.AddWindow(w2)

	server.AddSession(session)

	sessions := []*tmux.Session{}
	sessions = append(sessions, &session)

	conf := tmux.Configuration{
		Server:        server,
		Sessions:      sessions,
		ActiveSession: nil}

	// Setup this configuration.
	err = conf.Apply()
	if err != nil {
		msg := fmt.Errorf("Can't apply prepared configuration: %s", err)
		fmt.Println(msg)
		return
	}

	// Attach to created session
	// err = session.AttachSession()
	// if err != nil {
	// 	msg := fmt.Errorf("Can't attached to created session: %s", err)
	// 	fmt.Println(msg)
	// 	return
	// }
}


func getSessionInfo(session string, settings map[string]interface{}) []string {
  for k, v := range settings {
    if k == session {
      return ConvertToStringSlice(v)
    }
  }
  return nil
}


func CreateTmuxWindows(name string, settings map[string]interface{}) {
  // var session map[string]interface{}

  for k := range settings {
    if k == name {
      // fmt.Println(v)
      // fmt.Println(settings[k])
      // fmt.Println(reflect.TypeOf(settings[k]).Kind())
      win := reflect.ValueOf(settings[k])

      // for i := 0; i < win.Len(); i++ {
      //   fmt.Printf("%s", win.Index(i))
      // }
    }
  }
}
