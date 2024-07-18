# Bridge-exe



    go run github.com/diego-all/bridge-exe@latest init --name Diego
    go run main.go init --name Diego
    


ejecute sin subir los cambios es decir sin el CLI y es obvio:

    root@pho3nix:/home/diegoall/FALCO/bridge-exe# go run github.com/diego-all/bridge-exe@latest init --name Diego
    go: downloading github.com/diego-all/bridge-exe v0.0.0-20240718232318-84f438b75737
    Hola, mundo


Quizas sin ser CLI un script en go si pueda servir...


    go run main.go init --name Diego

root@pho3nix:/home/diegoall/FALCO/bridge-exe# go run main.go init --name Diego
main.go:3:8: no required module provides package github.com/diego-all/bridge-exe/cmd: go.mod file not found in current directory or any parent directory; see 'go help modules'
root@pho3nix:/home/diegoall/FALCO/bridge-exe# 

Obviamente no va a dar por que requiere el modulo.



