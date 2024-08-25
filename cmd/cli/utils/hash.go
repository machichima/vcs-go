// hash the serialized blob file
package utils

import (
	"crypto/sha1"
    "encoding/hex"
	"io"
)

// Hash the blob with sha1
func HashBlob(blob []byte) (string, error) {
    // Read the file to be hashed
    h := sha1.New()

    _, err := io.WriteString(h, string(blob))
    if err != nil {
        return "", err
    }
    hashStr := hex.EncodeToString(h.Sum(nil))
    return hashStr, nil
}

