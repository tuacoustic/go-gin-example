package commonConstants

import "time"

func NotNullBodyMsg() string {
	return "The body is not empty"
}

func JwtConstants() JwtKeyStructs {
	return JwtKeyStructs{
		PublicKey:  "middlewares/public.key",
		PrivateKey: "middlewares/private.key",
		Exp:        time.Now().Add(time.Minute * 1).Unix(),
	}
}
