package rpcat

import "os/exec"

// Cmd struct represents the output of cat command.
type Cmd struct {
	Output string
}

// Cat method, runs the cat command with the args
// provided by the client, and fills the cmd struct
// with the response.
func (c *Cmd) Cat(filename string, cout *Cmd) error {
	var std []byte
	cmd := exec.Command("cat", filename)
	std, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	cout.Output = string(std)
	return nil
}
