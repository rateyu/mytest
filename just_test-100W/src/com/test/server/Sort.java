package com.test.server;

/**
 * myu
 * 2019/2/14 22:07
 **/
public class Sort {

    public static void main(String[] args) {
        int[] a = {5, 8, 4, 3, 1, 9, 7,0};
        int l = a.length;
        //插入
        for (int i = 1; i < l; i++) {
            int t = a[i];
            int j = i - 1;
            for (; j >=0 ; j--) {
                if (a[j]>t)
                {
                    a[j + 1] = a[j];
                }
                else
                {
                    break;
                }
            }
            a[j+1] = t;
        }
        for (int i = 0; i < l; i++) {
            System.out.print(a[i]+",");
        }
    }
}
