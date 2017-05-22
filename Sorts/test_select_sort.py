# -*- coding: utf-8 -*-


from unittest import TestCase
from select_sort import select_sort


class SelectSortTestCase(TestCase):
    def setUp(self):
        self.sort = select_sort

    def test_quick_sort_one_list(self):
        lst = [1]

        self.sort(lst)

        self.assertEqual(lst, [1])


    def test_quick_sort_two_list(self):
        lst = [4, 6]

        self.sort(lst)

        self.assertEqual(lst, [4, 6])

    def test_quick_sort_two_list_2(self):
        lst = [6, 4]

        self.sort(lst)

        self.assertEqual(lst, [4, 6])


    def test_quick_sort_list_3(self):
        lst = [6, 3, 2, 5, 4]

        self.sort(lst)

        self.assertEqual(lst, [2,3,4,5,6])
