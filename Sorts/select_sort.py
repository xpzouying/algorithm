# -*- coding: utf-8 -*-


def select_sort(L):

    for i in range(len(L)):

        min = i
        
        for j in range(i+1, len(L)):
            if L[min] > L[j]:
                min = j

        L[min], L[i] = L[i], L[min]
