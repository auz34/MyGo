using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace WordDist3
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            var words = new string[] { "a", "a", "b", "b" };
            Console.WriteLine(s.ShortestWordDistance(words, "a", "b"));

            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int ShortestWordDistance(string[] words, string word1, string word2)
        {
            var word1Indices = new List<int>();
            var word2Indices = new List<int>();
            for (var i = 0; i < words.Length; i++)
            {
                if (words[i] == word1)
                {
                    word1Indices.Add(i);
                }

                if (words[i] == word2)
                {
                    word2Indices.Add(i);
                }
            }

            if (!word2Indices.Any() || !word2Indices.Any())
            {
                return -1;
            }

            var ptr1 = 0;
            var ptr2 = 0;
            var dist = Int32.MaxValue; 

            while (ptr1 != word1Indices.Count && ptr2 != word2Indices.Count)
            {
                var curDist = Math.Abs(word1Indices[ptr1] - word2Indices[ptr2]);

                if (curDist != 0 && curDist < dist)
                {
                    dist = curDist;
                }

                if (word1Indices[ptr1] < word2Indices[ptr2])
                {
                    ptr1++;
                }
                else
                {
                    ptr2++;
                }
            }

            return dist;
        }
    }
}
