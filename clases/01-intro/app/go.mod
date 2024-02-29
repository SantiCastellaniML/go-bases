module go-bases/intro/app

go 1.21.7

//go get github.com/fatih/color me trajo estas dependencias:

//cuando hago go mod tidy me limpia el archivo este y las indirectas pasan a ser directas.

//estas dependencias van a una carpeta que se llama /pkg/mod, es global.

//go.sum genera hashes a todos los packages. De esta forma asegura que la dependencia con la que estamos trabajando es realmente a la que referenciamos en el package mod.
//se utiliza para verificar seguridad e integridad de las referencias

require github.com/fatih/color v1.16.0

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.14.0 // indirect
)


//Go get -u all actualiza todas las extensiones