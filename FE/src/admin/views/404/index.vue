<template lang="pug">
.admin-content
  div.wrapper
    img(src="~common/assets/images/404.jpg")
    h1 对不起，你要访问的页面不存在
</template>

<script>
export default {
  data() {
    return {};
  },
};
</script>

<style lang="scss" scoped>
  .wrapper {
    text-align: center;
    margin-top: 1rem;

    h1 {
      margin-top: 0.12rem;
    }
  }
</style>
