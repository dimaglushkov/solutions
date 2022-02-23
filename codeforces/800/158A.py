n, k = [int(i) for i in input().split(' ')]
places = input().split(' ')
print(sum([1 for i in places if int(i) >= int(places[k - 1]) and int(i) > 0]))
