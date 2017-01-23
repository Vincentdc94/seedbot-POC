package botnet

import "net"
import "crypto/sha1"
import "encoding/hex"
import "os"

//MakeHash - Make a hash from a string and return it as string
func MakeHash(text string) string {
	hashMaker := sha1.New()
	hashMaker.Write([]byte(text))

	hash := hashMaker.Sum(nil)
	return hex.EncodeToString(hash)
}

//GetUniqueID -  get MAC adress as unique identifier for the running bot and convert it to hash
func GetUniqueID() string {

	interfaces, _ := net.Interfaces()
	hardwareAddr := interfaces[0].HardwareAddr.String()

	uniqueIdentifier := MakeHash(hardwareAddr)

	return uniqueIdentifier

}

//GetName - get computername of the bot for easier identification
func GetName() string {
	hostname, hostnameErr := os.Hostname()

	if hostnameErr != nil {
		return "Undefined Botname"
	}

	return hostname
}
