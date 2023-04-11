from functools import lru_cache

class Implementation:
    def __init__(self, n, m): # constructor
        self._n = n
        self._m = m
    
    @lru_cache(maxsize=None)
    def __waysToAttendClasses(self, n:int, m:int, k:int)->int:
        if (n==0):
            return 1
        
        miss = 0
        nomiss = 0
        if (k<m-1):
            miss = self.__waysToAttendClasses(n-1, m, k+1)
        nomiss = self.__waysToAttendClasses(n-1, m, 0)
        return miss+nomiss

    def __probToMissGraduationCeremony(self)->int: # private func
        return self.__waysToAttendClasses(self._n-1, self._m, 1)


    def __noOfWaysToAttendClasses(self)->int:
        return self.__waysToAttendClasses(self._n, self._m, 0)

    def __showInp(self):
        print(f"No. of academics days: {self._n}\nNo. of consecutive days allowed to miss the classes: {self._m}")

    def result(self)->str:
        a = self.__noOfWaysToAttendClasses()
        b = self.__probToMissGraduationCeremony()
        self.__showInp()
        res = f"Result = (number of ways to attend classes over N days/probability that you will miss your graduation ceremony)\n\t = {b}/{a}"
        return res
