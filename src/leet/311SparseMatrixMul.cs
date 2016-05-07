using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SparseMatirx
{
    class Program
    {
        static void Main(string[] args)
        {
            /*int[,] A = new int[,] { { 1, 0, 0 }, { -1, 0, 3 } };
            int[,] B = new int[,] { { 7, 0, 0 }, { 0, 0, 0 }, { 0, 0, 1 } };*/

            /*int[,] A = new int[,] { {1,-5} };
            int[,] B = new int[,] { { 12 }, { -1 } };*/

            int[,] A = new int[,] { { 1, 2, 3 }, { 4, 5, 6 } };
            int[,] B = new int[,] { { 7, 8, 9 }, { 10, 11, 12 }, { 13, 14, 15 } };




            Console.WriteLine("Matrix A:");
            printMatrix(A);
            Console.WriteLine("Matrix B:");
            printMatrix(B);

            var s = new Solution();
            Console.WriteLine("Mul:");
            printMatrix(s.Multiply(A, B));
            


            Console.ReadLine();
        }

        static void printMatrix(int[,] A)
        {
            for (int i = 0; i < A.GetLength(0); i++)
            {
                for (int j = 0; j < A.GetLength(1); j++)
                {
                    Console.Write(" {0} ", A[i, j]);
                }

                Console.WriteLine();
            }
        }
    }

    public class Solution
    {
        private int[,] VectorMultiply(int[,] A, int[,] B)
        {
            var m = A.GetLength(1);
            var result = new int[1, 1];
            for (int i = 0; i < m; i++)
            {
                result[0, 0] += A[0, i] * B[i, 0];
            }

            return result;
        }

        public HashSet<int> FindZeroRows(int[,] A)
        {
            var result = new HashSet<int>();
            var n = A.GetLength(0);
            var m = A.GetLength(1);

            for (int i = 0; i < n; i++)
            {
                bool allZero = true;
                for (int j = 0; j < m; j++)
                {
                    if (A[i, j] != 0)
                    {
                        allZero = false;
                        break;
                    }
                }
                if (allZero)
                {
                    result.Add(i);
                }
            }

            return result;
        }

        private HashSet<int> FindZeroColumns(int[,] A)
        {
            var result = new HashSet<int>();
            var n = A.GetLength(0);
            var m = A.GetLength(1);

            for (int j = 0; j < m; j++) 
            {
                bool allZero = true;
                for (int i = 0; i < n; i++)
                {
                    if (A[i, j] != 0)
                    {
                        allZero = false;
                        break;
                    }
                }
                if (allZero)
                {
                    result.Add(j);
                }
            }

            return result;
        }

        public int[,] Multiply(int[,] A, int[,] B)
        {
            var n = A.GetLength(0);
            var m = A.GetLength(1);

            if (n==1)
            {
                return VectorMultiply(A, B);
            }

            var result = new int[n, m];
            var zeroRows = FindZeroRows(A);
            var zeroColumns = FindZeroColumns(B);
            for (int i = 0; i < n; i++)
            {
                if (zeroRows.Contains(i))
                {
                    continue;
                }

                for (int j = 0; j < m; j++)
                {
                    if (zeroColumns.Contains(j))
                    {
                        continue;
                    }

                    for (int k = 0; k < m; k++)
                    {
                        var a = A[i, k];
                        var b = B[k, j];
                        result[i, j] += a * b;
                    }
                }
            }
            return result;
        }
    }
}
