using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace NestedInteger
{
    // This is the interface that allows for creating nested lists.
    // You should not implement it, or speculate about its implementation
    public interface NestedInteger
    {

        // @return true if this NestedInteger holds a single integer, rather than a nested list.
        bool IsInteger();

        // @return the single integer that this NestedInteger holds, if it holds a single integer
        // Return null if this NestedInteger holds a nested list
        int GetInteger();

        // @return the nested list that this NestedInteger holds, if it holds a nested list
        // Return null if this NestedInteger holds a single integer
        IList<NestedInteger> GetList();
    }

    public class NestedImpl: NestedInteger
    {
        private object content;

        public NestedImpl(object content)
        {
            this.content = content;
        }

        public bool IsInteger()
        {
            return !(this.content is IList<NestedInteger>);
        }

        public int GetInteger()
        {
            return (int)this.content;
        }

        public IList<NestedInteger> GetList()
        {
            return (IList<NestedInteger>)this.content;
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            var solution = new Solution();
            var list1 = new List<NestedInteger>() { new NestedImpl(1), new NestedImpl(1) };
            var case1list = new List<NestedInteger>()
                                {
                                    new NestedImpl(list1),
                                    new NestedImpl(2),
                                    new NestedImpl(list1)
                                };


            Console.WriteLine("{0}", solution.DepthSum(case1list));

            var inner3 = new List<NestedInteger>() { new NestedImpl(6) };
            var inner2 = new List<NestedInteger>() { new NestedImpl(4), new NestedImpl(inner3) };
            var case2list = new List<NestedInteger>() { new NestedImpl(1), new NestedImpl(inner2) };

            Console.WriteLine("{0}", solution.DepthSum(case2list));

            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int DepthSumInternal(IList<NestedInteger> nestedList, int level = 1)
        {
            var sum = 0;
            foreach (var item in nestedList)
            {
                if (item.IsInteger())
                {
                    sum += level * item.GetInteger();
                }
                else
                {
                    sum += this.DepthSumInternal(item.GetList(), level + 1);
                }
            }

            return sum;
        }
        public int DepthSum(IList<NestedInteger> nestedList)
        {
            return this.DepthSumInternal(nestedList);
        }
    }
}
