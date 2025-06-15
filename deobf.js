const parser = require('@babel/parser');
const traverse = require('@babel/traverse').default;
const generate = require('@babel/generator').default;
const fs = require('fs');
const vm = require('vm');
const t = require('@babel/types');

const code = fs.readFileSync('in.js', 'utf8');

const ast = parser.parse(code, {sourceType: 'script'});
let arrayname = "kianfr"
let vmcode = ""
let replaceFuncName = ""
traverse(ast, {
    VariableDeclarator(path) {
        if (path.node.init?.type === "ArrayExpression" && path.node.init.elements.length > 1000) {
            vmcode += generate(path.node, {}, code).code + ";\n";
            arrayname = path.node.id.name;
        } else if (
            path.node.init?.type === "FunctionExpression" &&
            path.node.init.params.length === 2
        ) {
            const dcode = generate(path.node, {}, code).code
            if (dcode.includes(arrayname)) {
                replaceFuncName = path.node.id.name;
            }
        }
    },
    ExpressionStatement(path) {
        const genned = generate(path.node, {}, code).code
        if (
            genned.includes("push")
            && genned.includes("shift")
            && genned.includes("function")
        ) {
            vmcode += generate(path.node, {}, code).code + "\n";
        }
    }

});

vmcode += arrayname + ";"
const array = vm.runInNewContext(vmcode, {});
console.log(replaceFuncName)


traverse(ast, {
    CallExpression(path) {
        const node = path.node
        if (
            node.callee.name !== replaceFuncName
            || node.arguments?.length !== 1
        ) return;

        const value = node.arguments[0].value
        path.replaceWith(t.stringLiteral(array[parseInt(value)]));
    }
})


const output = generate(ast, {}, code);
fs.writeFileSync('out.js', output.code);

console.log('Deobfuscation complete. Output saved to out.js');