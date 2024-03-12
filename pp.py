def go_to_rome(n):
    result = ''
    for arabic, roman in zip((1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1),
                             'M     CM   D    CD   C    XC  L   XL  X   IX V  IV I'.split()):
        result += n // arabic * roman
        n %= arabic
        #print('({}) {} => {}'.format(roman, n, result))
    return result



def go_to_rome2(n):
    result = ''
    for arabic, roman in ((1000, 'М'), (900, 'CM'), (500, 'D'), (400, 'CD'), (100, 'C'), (90, 'XC'), (50, 'L'), (40, 'XL'), (10, 'X'), (9, 'IX'), (5, 'V'), (4, 'IV'), (1, 'I')):
            result += n // arabic * roman
            n %= arabic
            #print('({}) {} => {}'.format(roman, n, result))
    return result

# XCIX = 99
# -10 100 -1 10
# 90 9 400 90
# III 

def go_to_rome3(n):
    dd = {
        "М":  100,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,}
    
    result = ''
    
    for key, values in dd.items():
        result += n // values * key
        print(n // values)
        n %= values
        
    
    return result

print(go_to_rome3(96))


def go_to_arab(n):
    rome_dict = {

    "I": 1,
    "IV": 4,
    "V": 5,
    "IX": 9,
    "X": 10,
    "L": 50,
    "XC": 90,
    "C": 100, 
    "CD": 400,	
    "D": 500,
    "CM": 900,	
    "M": 1000,

    }

    # XCVI = 96

    rome_num = n

    result = 0

    for y in range(len(rome_num)):                # проходимся циклом for по каждой римской цифре
        for ro, ar in rome_dict.items():          # проходися циклом for по словарю цифр пар: ключ: римские - значение: арабские
            if rome_num[y] == ro:                 # находим совпадение цифры из словаря
                if y == len(rome_num) - 1:        # проверяем не песледняя ли это римская цифра, если последняя, то просто приплюсовываем ее к результату и завершаем цикл
                    result += ar                  
                else:                             # если это не последняя римская цифра, то выполняем условия дальше 
                    if rome_num[y+1] in ["V", "X", "C", "D", "M"]: # проверям условие если следующая за проверяемой римской цифрой римская цифра из числа "V", "X", "C", "D", "M" 
                        ar = -ar                  # если это так, то отнять из результата
                        result += ar
                    else:                         # если ранее прописанные условия не выполенны, то просто прибавляем к результату
                        result += ar

    return result




def go_to_arab2(n):
    rome_dict = {

    "I": 1,
    "IV": 4,
    "V": 5,
    "IX": 9,
    "X": 10,
    "XL": 40,
    "L": 50,
    "XC": 90,
    "C": 100, 
    "CD": 400,	
    "D": 500,
    "CM": 900,	
    "M": 1000,

    }

    # XCVI = 96

    rome_num = n

    result = 0
    
    for y in range(len(rome_num)):                # проходимся циклом for по каждой римской цифре
        for ro, ar in rome_dict.items():          # проходися циклом for по словарю цифр пар: ключ: римские - значение: арабские
            if rome_num[y] == ro:                 # находим совпадение цифры из словаря
                if y == len(rome_num) - 1:        # проверяем не песледняя ли это римская цифра, если последняя, то просто приплюсовываем ее к результату и завершаем цикл
                    result += ar          
                else:                             # если это не последняя римская цифра, то выполняем условия дальше 
                    if rome_num[y] == "I" and rome_num[y+1] == "V":
                        ar = -ar
                        result += ar
                        break
                    if rome_num[y] == "I" and rome_num[y+1] == "X":
                        ar = -ar
                        result += ar
                        break
                    if rome_num[y] == "X" and rome_num[y+1] == "C":
                        ar = -ar
                        result += ar
                        break
                    if rome_num[y] == "X" and rome_num[y+1] == "L":
                        ar = -ar
                        result += ar
                        break
                    if rome_num[y] == "C" and rome_num[y+1] == "D":
                        ar = -ar
                        result += ar
                        break
                    if rome_num[y] == "C" and rome_num[y+1] == "M":
                        ar = -ar
                        result += ar
                        break
                    else:                         # если ранее прописанные условия не выполенны, то просто прибавляем к результату
                        result += ar
        
    return result



