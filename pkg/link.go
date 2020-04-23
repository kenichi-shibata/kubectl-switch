package pkg

import (
	"fmt"
	"os"
)

// Wrapper for softlinking kubectl-vx.x.x to kubectl
func SoftlinkKubectl(kubectlVersion, kubectl string) error {
	if _, err := os.Lstat(kubectl); err == nil {
		if err := os.Remove(kubectl); err != nil {
			return fmt.Errorf("failed to unlink: %+v", err)
		}
	} else if os.IsNotExist(err) {
		return fmt.Errorf("failed to check symlink: %+v", err)
	}
	err := os.Symlink(kubectlVersion, kubectl)
	if err != nil {
		return err
	}
	return nil
}
