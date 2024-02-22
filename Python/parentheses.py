class Skeleton:
    def isValid(self, question_string):
        # Create empty stack 
            
        stack = []

        for c in question_string:
        
            # Check the character at the i-th position.
            # If opening bracket -> add to stack.
            if c in "([{":
                stack.append(c)

            # Terminates function if stack is empty
            elif c in ")]}":
                if len(stack) == 0:
                    return False

                
                last = stack[-1]

                # If last closing bracket doesn't match current opening bracket -> false
                if (c == ')' and last != '(') or (c == ']' and last != '[') or (c == '}' and last != '{'):
                    return False

                # Pop last element from stack
                stack.pop()
        return len(stack) == 0

skeleton = Skeleton()
#string test case here
str = "({[])"
result = skeleton.isValid(str)
print(result)