package main

import (
	// "bytes"
	"io/fs"
)

type fileStruct struct {
	Path string
	Mode fs.FileMode
}

const asciiArt string = `
  /&&&&&\    
 &&/   \&&   .__                       __
 &&.. ..&&   |  |     ____     ____   |  | __   ____   _______
&&&&&&&&&&&  |  |    /  _ \  _/ ___\  |  |/ / _/ __ \  \_  __ \
&&&&' '&&&&  |  |__ (  <_> ) \  \___  |    <  \  ___/   |  | \/
&&&&& &&&&&  |____/  \____/   \___  > |__|_ \  \___  >  |__|    
&&&&&&&&&&&                       \/       \/      \/           
`;

const version string = "beta 0.1"

const signature string = "locked by locker " + version + "\n";

const configPath string = "./config.json";

const vaultPath string = "./vault.json";

// how many times you want to verify data, 1=off
const verifyData int = 2;

var salt = "DEFAULTPASSWORD";
