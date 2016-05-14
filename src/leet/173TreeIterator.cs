using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TreeIterator
{
    using System.ComponentModel.Design.Serialization;
    using System.Runtime.CompilerServices;

    class Program
    {
        static void Main(string[] args)
        {
            var node5 = new TreeNode(5);
            var node7 = new TreeNode(7);
            var node9 = new TreeNode(9);
            var node11 = new TreeNode(11);

            var node6 = new TreeNode(6) {left = node5, right = node7};
            var node10 = new TreeNode(10) { left = node9, right = node11 };

            var node8 = new TreeNode(8) { left = node6, right = node10};

            Console.WriteLine("Use case #1");
            BSTIterator i = new BSTIterator(node8);
            while (i.HasNext())
            {
                Console.WriteLine(i.Next()); 
            }

            Console.WriteLine("Use case #2");
            var node22 = new TreeNode(22);
            var node21 = new TreeNode(21) {right = node22};
            BSTIterator i2 = new BSTIterator(node21);
            while (i2.HasNext())
            {
                Console.WriteLine(i2.Next());
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
 

    public class BSTIterator
    {

        private Stack<TreeNode> stack = new Stack<TreeNode>();

        private TreeNode nextNode;

        public BSTIterator(TreeNode root)
        {
            this.MoveToLeft(root);
        }

        private bool IsLeaf(TreeNode node)
        {
            return node.left == null && node.right == null;
        }

        private void MoveToLeft(TreeNode node)
        {
            this.nextNode = node;
            if (node != null && !this.IsLeaf(node))
            {
                this.stack.Push(node);
                this.MoveToLeft(node.left);
            }
        }

        /** @return whether we have a next smallest number */
        public bool HasNext()
        {
            return this.nextNode != null || this.stack.Any();
        }

        /** @return the next smallest number */
        public int Next()
        {
            if (this.nextNode != null)
            {
                var result = this.nextNode.val;
                this.nextNode = null;
                return result;
            }

            var subRoot = this.stack.Pop();
            this.MoveToLeft(subRoot.right);
            return subRoot.val;
        }
    }
}
