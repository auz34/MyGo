using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ZigZagTraverse
{
    using System.ComponentModel;

    class Program
    {
        static void Main(string[] args)
        {
            var node5 = new TreeNode(5);
            var node7 = new TreeNode(7);
            var node9 = new TreeNode(9);
            var node11 = new TreeNode(11);

            var node6 = new TreeNode(6) { left = node5, right = node7 };
            var node10 = new TreeNode(10) { left = node9, right = node11 };

            var node8 = new TreeNode(8) { left = node6, right = node10 };

            var solution = new Solution();
            var res = solution.ZigzagLevelOrder(node8);

            foreach (var level in res)
            {
                foreach (var num in level)
                {
                    Console.Write(" {0} ", num);
                }

                Console.WriteLine();
            }

            Console.ReadLine();

        }
    }


    public class TreeNode {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
      }

    public class LevelCollection
    {
        private string type;
        private object internalCollection;
        public LevelCollection(string type)
        {
            this.type = type;
            if (type == "stack")
            {
                internalCollection = new Stack<int>();
            }
            else
            {
                internalCollection = new List<int>();
            }
        }

        public void Add(int value)
        {
            if (this.type == "stack")
            {
                ((Stack<int>)this.internalCollection).Push(value);
            }
            else
            {
                ((List<int>)this.internalCollection).Add(value);
            }
        }

        public List<int> GetList()
        {
            if (this.type == "stack")
            {
                return ((Stack<int>)this.internalCollection).ToList();
            }
            else
            {
                return (List<int>)this.internalCollection;
            }
        }

    }
 
    public class Solution
    {
        private void internalTraverse(TreeNode node, int level, List<LevelCollection> result)
        {
            if (node == null)
            {
                return;
            }

            if (result.Count <= level)
            {
                result.Add(new LevelCollection((level % 2 == 0) ? "list" : "stack"));
            }

            result[level].Add(node.val);
            this.internalTraverse(node.left, level + 1, result);
            this.internalTraverse(node.right, level + 1, result);
        }
        public IList<IList<int>> ZigzagLevelOrder(TreeNode root)
        {
            var result = new List<LevelCollection>();
            this.internalTraverse(root, 0, result);

            IList<IList<int>> answer = new List<IList<int>>();
            foreach (var levelCollection in result)
            {
                answer.Add(levelCollection.GetList());
            }

            return answer;
        }
    }
}
