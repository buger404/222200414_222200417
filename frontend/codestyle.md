# 前端代码规范

**来源**: [Vue 官方风格指南](https://vuejs.org/style-guide/), [TypeScript 官方风格指南](https://github.com/microsoft/TypeScript/wiki/Coding-guidelines), [Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript)

## 1. 通用规范
- **缩进**：使用 2 个空格缩进。
- **分号**：每行末尾必须使用分号。
- **引号**：使用单引号 `''`，除非需要转义。
- **文件名**：使用 `kebab-case` 命名组件和文件。
- **注释**：优先使用 `//` 进行单行注释，`/** */` 进行块注释。

## 2. TypeScript 规范
- **类型声明**：优先使用接口 `interface` 而不是类型别名 `type`。
- **空值检查**：开启 `strictNullChecks` 并始终显式处理 `null` 和 `undefined`。
- **泛型命名**：泛型参数使用简短有意义的名称，如 `T` 表示 `Type`，`K` 表示 `Key`。
- **函数返回类型**：始终显式声明函数的返回类型。

```ts
interface User {
  name: string;
  age: number;
}

function getUserName(user: User): string {
  return user.name;
}
```

## 3. Vue 规范

- **组件命名**：组件名使用 `PascalCase`，标签使用 `kebab-case`。
- **Prop**：始终为 `prop` 定义类型。
- **样式作用域**：尽量使用 `scoped` 或 `module` 来限制样式作用域。
- **事件命名**：自定义事件应使用 `kebab-case`，并以 `on` 作为监听前缀。

```vue
<template>
  <MyComponent @my-event="handleEvent" />
</template>

<script lang="ts">
export default defineComponent({
  props: {
    title: {
      type: String,
      required: true
    }
  },
  methods: {
    handleEvent() {
      console.log('Event triggered');
    }
  }
});
</script>

<style scoped>
/* 样式只在当前组件作用域内生效 */
</style>
```

## 4. 推荐工具

- 使用 [ESLint](https://eslint.org/) 和 [Prettier](https://prettier.io/) 进行代码风格检查和格式化。
- Vue 项目建议使用 Vue ESLint 插件 进行代码校验。