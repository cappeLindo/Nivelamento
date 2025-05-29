a = 0
b = 1

num = int(input('Digite um número: '))

while num > 1:
    print(b)
    c = a
    a = b
    b = a + c
    num -= 1

print('acabou:)')



