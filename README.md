0 define jogadores
0.1 define equipe a iniciar

1 serve
    let
        confirma
        voltar p/ 1 
    dentro
        confirma
        ir p/ 2
    fora
        confirma
        aculuma 1 falta
        voltar p/ 1

2 retorna serviço
    não toca na bola
        confirma
        atualiza placar
    dentro
        confirma
        ir p/ 3
    erro: Jogou fora
        confirma
        atualiza placar
    erro: Jogou na rede
        confirma
        atualiza placar


3 retorna comum
    não toca na bola
        confirma
        atualiza placar
    dentro
        confirma
        ir p/ 3
    erro: Jogou fora
        confirma
        atualiza placar
    erro: Jogou na rede
        confirma
        atualiza placar
