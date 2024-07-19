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



root@pho3nix:/home/diegoall/FALCO/bridge-exe# go run main.go init --name Diego
Hola Diego
root@pho3nix:/home/diegoall/FALCO/bridge-exe# go run github.com/diego-all/bridge-exe@latest init --name Diego
Hola, mundo
root@pho3nix:/home/diegoall/FALCO/bridge-exe# git status
On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean


HALLAZGO A PESAR DE QUE EL REMOTO ESTA ACTUALIZADO, SIGUE BUSCANDO UNA VERSION ANTIGUA.


segun chat gpt 8 No use fregierr

    git tag -a v0.1.1 -m "Nueva versión con el comando init"
    git push origin v0.1.1
    go run github.com/diego-all/bridge-exe@v0.1.1 init --name Diego


root@pho3nix:/home/diegoall/FALCO/bridge-exe# go run github.com/diego-all/bridge-exe@v0.1.1 init --name Diego
go: downloading github.com/diego-all/bridge-exe v0.1.1
Hola Diego



Funcionma bien desde un equipo local ejecutando desde repo remoto.


Ahora a hacerlo a travez de python

    root@pho3nix:/home/diegoall/FALCO/falco-workshop# curl -sL http://34.27.180.215/sitio/vuelta.txt | python3
    Salida estándar:
    Hola Diego

    Salida de error:


Si funciona con python, pero requiere apuntar a una version especifica.

Que sera entonces con el workshop?