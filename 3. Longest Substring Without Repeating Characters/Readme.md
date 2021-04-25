# Idea

就暴力法，從頭到尾找看看不重複的字串最大是多少

但實作上面要想到用hashMap紀錄字元出現的最後位置。

當出現重複字元時就可以依hashMap裡記錄的位置進行更新新的起始位置





# Time Complexity

剛好把string從頭到尾走完所以是 O(N)