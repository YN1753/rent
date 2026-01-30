<!-- <template>
  <div class="login-container">
    <div class="login-switch">
      <router-link class="switch-item active" to="/login/account">用户名登录</router-link>
      <router-link class="switch-item" to="/login/email">邮箱登录</router-link>
    </div>
    <van-form @submit="onAccountSubmit">
      <van-field v-model="accountForm.username" name="username" label="用户名" placeholder="请输入用户名" :rules="[{ required: true, message: '请填写用户名' }]" />
      <van-field v-model="accountForm.password" :type="showPassword ? 'text' : 'password'" name="password" label="密码" placeholder="请输入密码" :rules="[{ required: true, message: '请填写密码' }]">
        <template #button>
          <van-icon :name="showPassword ? 'eye-o' : 'closed-eye'" size="18" @click="togglePassword" />
        </template>
      </van-field>
      <div class="remember">
        <van-checkbox v-model="remember">记住用户名</van-checkbox>
      </div>
      <div style="margin: 16px;">
        <van-button round block type="primary" native-type="submit">登录</van-button>
      </div>
    </van-form>
    <div class="footer-links">
      <router-link to="/register">注册账号</router-link>
      <span>|</span>
      <router-link to="/forget">忘记密码?</router-link>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/store/modules/user'
const router = useRouter()
const userStore = useUserStore()
const accountForm = ref({ username: '', password: '' })
const showPassword = ref(false)
const remember = ref(false)
const togglePassword = () => { showPassword.value = !showPassword.value }
const onAccountSubmit = async (values) => {
  const success = await userStore.login({ type: 'account', username: values.username, password: values.password })
  if (success) {
    if (remember.value) localStorage.setItem('remembered_account', values.username)
    else localStorage.removeItem('remembered_account')
    router.push('/home')
  }
}
</script>
<style scoped>
.login-container { max-width: 360px; margin: 40px auto; padding: 0 16px; }
.login-switch { display: flex; gap: 12px; margin-bottom: 16px; }
.switch-item { text-decoration: none; color: #666; padding: 8px 12px; border-radius: 16px; background: #f2f3f5; }
.switch-item.active { color: #fff; background: #1989fa; }
.remember { margin: 8px 0; text-align: right; }
.footer-links { margin-top: 12px; text-align: center; color: #1989fa; }
.footer-links a { color: #1989fa; margin: 0 6px; }
</style> -->

<!-- version 2 -->
 <template>
  <!-- 包裹 Modal 背景 -->
  <div class="modal-overlay" @click="close">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h3>用户名登录</h3>
        <van-icon name="close" size="18" @click="close" />
      </div>

      <!-- 表单内容 -->
      <div class="login-switch">
        <span class="switch-item active">用户名登录</span>
        <span class="switch-item" @click="switchToEmail">邮箱登录</span>
      </div>

      <van-form @submit="onAccountSubmit">
        <van-field
          v-model="accountForm.username"
          name="username"
          label="用户名"
          placeholder="请输入用户名"
          :rules="[{ required: true, message: '请填写用户名' }]"
        />
        <van-field
          v-model="accountForm.password"
          :type="showPassword ? 'text' : 'password'"
          name="password"
          label="密码"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请填写密码' }]"
        >
          <template #button>
            <van-icon
              :name="showPassword ? 'eye-o' : 'closed-eye'"
              size="18"
              @click="togglePassword"
            />
          </template>
        </van-field>
        <div class="remember">
          <van-checkbox v-model="remember">记住用户名</van-checkbox>
        </div>
        <div style="margin: 16px;">
          <van-button round block type="primary" native-type="submit">
            登录
          </van-button>
        </div>
      </van-form>

      <div class="footer">
        <p>
          <span @click="openRegister">注册账号</span> |
          <router-link to="/forget">忘记密码?</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/store/modules/user'
import { Icon } from 'vant'

const router = useRouter()
const userStore = useUserStore()

const accountForm = ref({ username: '', password: '' })
const showPassword = ref(false)
const remember = ref(false)


const togglePassword = () => {
  showPassword.value = !showPassword.value
}

// 登录提交
const onAccountSubmit = async (values) => {
  const success = await userStore.login({
    type: 'account',
    username: values.username,
    password: values.password
  })
  if (success) {
    if (remember.value) localStorage.setItem('remembered_account', values.username)
    else localStorage.removeItem('remembered_account')
    // 弹窗登录成功后不需要 router.push，直接关闭即可
    close()
  }
}

// 关闭页面：作为路由页面，返回上一页或回到首页
const close = () => {
  try {
    router.go(-1)
  } catch (e) {
    router.push('/home')
  }
}

// 切换到邮箱登录：直接跳路由
const switchToEmail = () => {
  router.push('/login/email')
}

// 打开注册页：直接跳路由
const openRegister = () => {
  router.push('/register')
}
</script>

<!-- 注册弹窗样式 -->
<style scoped>
/* Modal 基础样式 */
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
.footer a,
.footer span {
  color: #1989fa;
  text-decoration: underline;
  cursor: pointer;
}

/* 登录切换样式 */
.login-switch {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
}
.switch-item {
  text-decoration: none;
  color: #666;
  padding: 8px 12px;
  border-radius: 16px;
  background: #f2f3f5;
  cursor: pointer;
}
.switch-item.active {
  color: #fff;
  background: #1989fa;
}
.remember {
  margin: 8px 0;
  text-align: right;
}
</style>
