package utils

import (
	"fmt"

	tmux "github.com/jubnzv/go-tmux"
)

func CreateServer(namespace string) (*tmux.Server, tmux.Session) {
	server := new(tmux.Server)

	// Check that "example" session already exists.
	exists, err := server.HasSession(namespace)
	if err != nil {
		msg := fmt.Errorf("Can't check '%s' session: %s", namespace, err)
		fmt.Println(msg)
		return nil, tmux.Session{}
	}

	if exists {
		// Sure, you can use KillSession here.
		fmt.Printf("Session '%s' already exists!", namespace)
		fmt.Println("Please stop it before running this demo.")
		return nil, tmux.Session{}
	}

	// Prepare configuration for a new session with some windows.
	session := tmux.Session{Name: fmt.Sprintf("bg-%s", namespace)}

	return server, session

	// w1 := tmux.Window{Name: "first", Id: 0, StartDirectory:}
	// w2 := tmux.Window{Name: "second", Id: 1}
	//
	// session.AddWindow(w1)
	// session.AddWindow(w2)
	//
	// server.AddSession(session)
	//
	// sessions := []*tmux.Session{}
	// sessions = append(sessions, &session)

	// conf := tmux.Configuration{
	// 	Server:        server,
	// 	Sessions:      sessions,
	// 	ActiveSession: nil}
	//
	// // Setup this configuration.
	// err = conf.Apply()
	// if err != nil {
	// 	msg := fmt.Errorf("Can't apply prepared configuration: %s", err)
	// 	fmt.Println(msg)
	// 	return
	// }

	// Attach to created session
	// err = session.AttachSession()
	// if err != nil {
	// 	msg := fmt.Errorf("Can't attached to created session: %s", err)
	// 	fmt.Println(msg)
	// 	return
	// }
}

func AddWindow(server *tmux.Server, session tmux.Session, directory string, cmd string, index int) {
	window := tmux.Window{Name: cmd, Id: index, StartDirectory: directory}

	session.AddWindow(window)

	server.AddSession(session)
}

func Start(server *tmux.Server, sessions []*tmux.Session) {
	conf := tmux.Configuration{
		Server:        server,
		Sessions:      sessions,
		ActiveSession: nil}

	// Setup this configuration.
	err := conf.Apply()
	if err != nil {
		msg := fmt.Errorf("Can't apply prepared configuration: %s", err)
		fmt.Println(msg)
		return
	}
}
