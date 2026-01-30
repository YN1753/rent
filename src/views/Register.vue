<!-- version 2 --
<!-- 注册页面 -->
<template>
  <div class="modal-overlay" @click="close">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h3>注册</h3>
        <van-icon name="close" size="18" @click="close" />
      </div>

      <van-form @submit="onSubmit">
        <van-field
          v-model="form.username"
          name="username"
          label="用户名"
          placeholder="请输入用户名"
          type="text"
          :rules="[{ required: true, message: '请填写用户名' }]"
        />
        <van-field
          v-model="form.email"
          name="email"
          label="邮箱"
          placeholder="请输入邮箱"
          :rules="[
            { required: true, message: '请填写邮箱' },
            { validator: isEmail, message: '邮箱格式错误' }
          ]"
        />
        <van-field
          v-model="form.code"
          name="code"
          label="验证码"
          placeholder="请输入验证码"
          :rules="[{ required: true, message: '请填写验证码' }]"
        >
          <template #button>
            <van-button size="small" type="primary" :disabled="countDown > 0" @click="sendCode">
              {{ countDown ? `${countDown}s后重发` : '获取验证码' }}
            </van-button>
          </template>
        </van-field>
        <van-field
          v-model="form.password"
          type="password"
          label="密码"
          placeholder="请输入6-20位密码"
          :rules="[
            { required: true, message: '请填写密码' },
            { pattern: /^.{6,20}$/, message: '密码长度为6-20位' }
          ]"
        />

        <div style="margin: 16px;">
          <van-button round block type="primary" native-type="submit">
            注册
          </van-button>
        </div>
      </van-form>

      <div class="footer">
        <p>已有账号？<router-link to="/login">去登录</router-link></p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store/modules/user';
import { sendCodeAPI, verifyCodeAPI } from '@/api/user';

const router = useRouter();
const userStore = useUserStore();

const form = ref({
  username: '',
  email: '',
  code: '',
  password: ''
});
const countDown = ref(0);

// 发送验证码
const sendCode = async () => {
  if (!isEmail(form.value.email)) {
    alert('请先填写正确的邮箱');
    return;
  }
  try {
    await sendCodeAPI({ email: form.value.email });
    startCountDown();
  } catch (e) {
    alert('验证码发送失败，请稍后重试');
  }
};

// 倒计时
const startCountDown = () => {
  countDown.value = 60;
  const timer = setInterval(() => {
    countDown.value--;
    if (countDown.value <= 0) clearInterval(timer);
  }, 1000);
};

// 注册提交
const onSubmit = async () => {
  // 先校验验证码
  try {
    await verifyCodeAPI({ email: form.value.email, code: form.value.code })
  } catch (e) {
    alert('验证码校验失败，请重新获取或重试')
    return
  }
  // 校验通过再注册
  const ok = await userStore.register(form.value)
  if (ok) {
    close()
  } else {
    alert('注册失败，请稍后重试')
  }
};

// 邮箱校验
const isEmail = (val) => /^\w+([-.+]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/.test(val);

// 关闭弹窗：作为路由页面，直接跳转回首页或上一个页面
const close = () => {
  try {
    router.go(-1);
  } catch (e) {
    router.push('/home');
  }
};

</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 200;
}
.modal-content {
  background: white;
  border-radius: 8px;
  width: 300px;
  max-width: 90%;
  padding: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
.footer {
  margin-top: 20px;
  text-align: center;
  font-size: 12px;
  color: #999;
}
.footer a {
  color: #1989fa;
  text-decoration: underline;
}
</style>

