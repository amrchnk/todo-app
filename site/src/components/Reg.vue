<template>
  <main>
    <div class="wrap">
      <div class="auth_form">
        <div class="fields">
          <div class="field">
            <h1>Регистрация</h1>
          </div>

          <div class="field">
            <label class="form-label">Имя пользователя</label>
            <input class="form-control" id="exampleInputEmail1" v-model="user.name">
          </div>

          <div class="field">
            <label class="form-label">Логин</label>
            <input class="form-control" id="exampleInputEmail1" v-model="user.username">
          </div>

          <div class="field">
            <label class="form-label">Пароль</label>
            <input class="form-control" id="exampleInputPassword" type="password" v-model="user.password">
          </div>
          <div class="others">
            <button  @click="reg">Зарегистрироваться</button>
            <router-link class="rl link" to="/">Назад ко входу</router-link>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: "Reg",
  data() {
    return {
      address: "http://localhost:8000",
      user: {
        name: "",
        username: "",
        password: ""
      },
      answer: {}
    }
  },
  methods:{
    async reg(){
      await this.$http.post(`${this.address}/auth/sign-up`, this.user) //по JSON RPC запросы идут только через POST
          .then((res) => res.json()).then((res) => (this.answer = res.result)); //ответ берем обязательно из res.result
      console.log(this.user);
      // if (this.user.name.length > 0 && this.user.username.length > 0 && this.user.password.length > 0) {
      //   this.$http.post(`${this.address}/auth/sign-up`, request).success(function(res) {
      //     console.log(res)
      //   }).error(function(err) {
      //     console.log(err);
      //   });
        // this.$http.post(`${this.address}/auth/sign-up`, request)
        //         .then((res) => res.json()).then((res) => (this.answer = res));
        // console.log(this.answer
    }
  }
}
</script>

<style scoped>
main{
  display: flex;
  justify-content: center;
  align-items: center;
}

.wrap{
  width: 80%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.auth_form{
  font-size: 1.3em;
  margin-top:20px;
  background-color: white;
  border-radius: 10px;
  border: 0.5px solid #E5E5E5;
  width: 32em;
}

h1{
  font-weight: 400;
}

.fields{
  padding: 40px 50px;
}
.field{
  width:100%;
  margin-bottom: 25px;
  clear:both;
  text-align:left;
  font-weight: 300;
}
.others{
  font-weight: 400;
  display: flex !important;
  align-items: center!important;
  justify-content: space-between;
  margin-top: 40px;
  color: #18175F;
}
button{
  font-weight: 500;
  font-size: 1em;
  padding: 5px 20px;
  background-color: #F08BC2;
  border: none;
  border-radius: 3px;
  color: #18175F;
}
.link{
  color: inherit;
  text-decoration: none;
}

</style>