def opcao():
    print('[1] Multiplicar números.')
    print('[2] Sair da aplicação.')

def multiplicar_numeros():
    n1 = float(input('Digite um número: '))
    n2 = float(input('Digite outro número: '))
    mult = n1 * n2
    print(37*'-')
    return f'A multiplicação de {n1} x {n2} = {mult}'
    


def continuar():
    cont = True
    while cont == True:
        opcao()
        opcoes = int(input('Digite uma opção: '))
        match opcoes:
            case 1:
                print(multiplicar_numeros())
                print(37*'-')
            case 2:
                print('Até mais :)')
                cont = False
                break
            case _:
                print('Opção inválida!')

continuar()


##ignore