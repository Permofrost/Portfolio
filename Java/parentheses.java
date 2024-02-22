import java.util.*;

public class parentheses {
    public boolean isValid(String question_String) {
        // Create empty stack 
        Stack<Character> stack = new Stack<Character>();

       
       
        for (int i = 0; i < question_String.length(); i++) {
            char c = question_String.charAt(i);

            // Check the character at the i-th position.
            // If opening bracket -> add to stack.
            if (c == '(' || c == '[' || c == '{') {
                stack.push(c);
            } 
            
            // Terminates function if stack is empty
            else {
                if (stack.isEmpty()) {
                    return false;
                }
            
            // If last closing bracket doesn't match current opening bracket -> false
                char last = stack.peek();
                if (
                    (c == ')' && last != '(') 
                    || (c == ']' && last != '[') 
                    || (c == '}' && last != '{')) {
                    return false;
                }
            // Pop last element from stack
                stack.pop();
            }
        }
        return stack.isEmpty();
}


    public static void main(String[] args) {
        
        parentheses p = new parentheses();    
        //Insert Test Case Here:
        String str = "({[]})";
        boolean result = p.isValid(str);
            
            System.out.println(result);
    }
}

/*
for _, c := range question_string {


    switch c {
    case '(', '[', '{':
        stack = append(stack, c)

    // If closing bracket -> first check if stack empty -> false if empty
    // Otherwise compare last opening bracket to current closing bracket
    case ')', ']', '}':

        // Terminates function if stack is empty
        if len(stack) == 0 {
            return false
        }

        // If last closing bracket doesn't match current opening bracket -> false
        last := stack[len(stack)-1]
        if (c == ')' && last != '(') || (c == ']' && last != '[') || (c == '}' && last != '{') {
            return false
        }
        // Pops last element from stack <----- ITERATION OCCURS HERE
        stack = stack[:len(stack)-1]
    }
}
return true
}


}
*/