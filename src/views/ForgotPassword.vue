<!-- <template>
  <div class="forget-container">
    <van-nav-bar title="找回密码" left-text="返回" @click-left="goBack" />

    <van-form @submit="onSubmit">
      <van-field
        v-model="form.email"
        name="email"
        label="注册邮箱"
        placeholder="请输入您的注册邮箱"
        :rules="[
          { required: true, message: '请填写邮箱' },
          { validator: isEmail, message: '邮箱格式错误' }
        ]"
      />
      <van-field
        v-model="form.code"
        center
        clearable
        name="code"
        label="验证码"
        placeholder="请输入验证码"
        :rules="[{ required: true, message: '请填写验证码' }]"
      >
        <template #button>
          <van-button size="small" :disabled="countDown > 0">
            {{ countDown ? `${countDown}s后重发` : '获取验证码' }}
          </van-button>
        </template>
      </van-field>
      <van-field
        v-model="form.newPassword"
        type="password"
        name="newPassword"
        label="新密码"
        placeholder="请输入6-20位新密码"
        :rules="[
          { required: true, message: '请填写新密码' },
          { pattern: /^.{6,20}$/, message: '密码长度为6-20位' }
        ]"
      />

      <div style="margin: 16px;">
        <van-button round block type="primary" native-type="submit">
          重置密码
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const form = ref({
  email: '',
  code: '',
  newPassword: ''
});
const countDown = ref(0);

const isEmail = (val) => /^\w+([-.+]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/.test(val);

// 发送验证码
const sendCode = () => {
  if (isEmail(form.value.email)) {
    startCountDown();
  } else {
    alert('请先输入正确的邮箱');
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

const onSubmit = async (values) => {
  console.log('重置密码信息:', values);
  alert('密码已重置，请返回登录');
  router.push('/login');
};

const goBack = () => router.go(-1);
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
 -->

<!-- version 2 -->
 <template>
  <div class="modal-overlay" @click="goBack">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h3>找回密码</h3>
        <van-icon name="close" size="18" @click="goBack" />
      </div>

      <van-form @submit="onSubmit">
        <van-field
          v-model="form.email"
          name="email"
          label="注册邮箱"
          placeholder="请输入您的注册邮箱"
          :rules="[
            { required: true, message: '请填写邮箱' },
            { validator: isEmail, message: '邮箱格式错误' }
          ]"
        />

        <van-field
          v-model="form.code"
          center
          clearable
          name="code"
          label="验证码"
          placeholder="请输入验证码"
          :rules="[{ required: true, message: '请填写验证码' }]"
        >
          <template #button>
            <van-button
              size="small"
              :disabled="countDown > 0"
              @click.prevent="sendCode"
            >
              {{ countDown ? ` $ {countDown}s后重发` : '获取验证码' }}
            </van-button>
          </template>
        </van-field>

        <van-field
          v-model="form.newPassword"
          type="password"
          name="newPassword"
          label="新密码"
          placeholder="请输入6-20位新密码"
          :rules="[
            { required: true, message: '请填写新密码' },
            { pattern: /^.{6,20} $ /, message: '密码长度为6-20位' }
          ]"
        />

        <div style="margin: 16px;">
          <van-button round block type="primary" native-type="submit">
            重置密码
          </van-button>
        </div>
      </van-form>

      <!-- 底部返回登录链接 -->
      <div class="footer">
        <p>
          <span @click="goBack">返回登录</span>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Icon, Button } from 'vant'

const router = useRouter()
const form = ref({
  email: '',
  code: '',
  newPassword: ''
})
const countDown = ref(0)

// 邮箱校验
const isEmail = (val) => /^\w+([-.+]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)* $ /.test(val)

// 发送验证码
const sendCode = () => {
  if (!isEmail(form.value.email)) {
    // vant 表单会自动校验，这里可加 toast 提示（可选）
    return
  }
  startCountDown()
}

// 倒计时
const startCountDown = () => {
  countDown.value = 60
  const timer = setInterval(() => {
    countDown.value--
    if (countDown.value <= 0) clearInterval(timer)
  }, 1000)
}

// 提交
const onSubmit = async (values) => {
  console.log('重置密码信息:', values)
  alert('密码已重置，请返回登录')
  goBack()
}

// 返回上一页（或首页）
const goBack = () => {
  try {
    router.go(-1)
  } catch (e) {
    router.push('/login/account')
  }
}
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
.modal-header h3 {
  font-size: 18px;
  margin: 0;
}

.footer {
  margin-top: 20px;
  text-align: center;
  font-size: 12px;
  color: #999;
}
.footer span {
  color: #1989fa;
  text-decoration: underline;
  cursor: pointer;
}
</style>