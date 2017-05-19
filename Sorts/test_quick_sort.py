# -*- coding: utf-8 -*-


from unittest import TestCase
from quick_sort import quicksort_by_wall


class QuickSortByWallTestCase(TestCase):
    def setUp(self):
        self.quick_sort = quicksort_by_wall

    def test_quick_sort_one_list(self):
        lst = [1]

        self.quick_sort(lst, 0, len(lst)-1)

        self.assertEqual(lst, [1])


    def test_quick_sort_two_list(self):
        lst = [4, 6]

        self.quick_sort(lst, 0, len(lst)-1)

        self.assertEqual(lst, [4, 6])

    def test_quick_sort_two_list_2(self):
        lst = [6, 4]

        self.quick_sort(lst, 0, len(lst)-1)

        self.assertEqual(lst, [4, 6])


    def test_quick_sort_list_3(self):
        lst = [6, 3, 2, 5, 4]

        self.quick_sort(lst, 0, len(lst)-1)

        self.assertEqual(lst, [2,3,4,5,6])
