<template>
  <main>
    <div class="wrap" v-if="!token">
      <div class="auth_form" >
        <div class="fields">
          <div class="field">
            <h1>Вход</h1>
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
            <button @click="auth">Войти</button>
            <router-link class="rl link" to="/reg">У меня нет аккаунта</router-link>
          </div>
        </div>
      </div>
    </div>
    <div class="wrap" v-if="token">
<!--      <h1>Привет, {{this.user.username}}!</h1>-->
      <div class="list_form">
        <div class="list_form_header">
          <input class="list_header" type="text" placeholder="Введите заголовок...">
          <hr>
        </div>
        <div class="add_items">
          <div class="plus">

          </div>
          <div class="add_items_input">
            <input class="add_input" type="text" placeholder="Новый пункт">
          </div>
          <hr>
        </div>
        <div class="list_items">

        </div>
      </div>
    </div>
  </main>
</template>

<script>

export default {
  name: "Main",
  data() {
    return {
      address: "http://localhost:8000",
      user: {
        username: '',
        password:''
      },
      answer:{},
      token:'',
      items:[]
    }
  },
  methods:{
    async auth(){
      if (this.user.username.length > 0 && this.user.password.length > 0) {
        try{
          await this.$http.post(`${this.address}/auth/sign-in`, this.user)
              .then((res) => res.json()).then((res) => (this.answer=res))
              .catch((res) => (this.answer=res["data"]));
          // console.log(this.answer["message"])
          if (this.answer["message"]===undefined && this.answer["token"]!=null){
            this.token=this.answer["token"]
            document.cookie = `username=${this.user.username}`;
            document.cookie = `password=${this.user.password}`;
            document.cookie=`token=${this.token}`
            // alert("Успешно")
          }
          else {
            alert("Ошибка в веденных данных")
          }
        }catch (e){
          console.log(e)
        }
      }
      else{
        alert("Поля формы не должны быть пустыми!")
      }
    },
    setToken() {
      // this.user.token = this.cookie.replace(/(?:(?:^|.*;\s*)token=\s*\s*([^;]*).*$)|^.*$/, "$1");
      // this.password=this.cookie.replace(/(?:(?:^|.*;\s*)full_name=\s*\s*([^;]*).*$)|^.*$/, "$1");
    },
  },
  mounted() {
    this.token
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

  .list_form{
    width: 50%;
    background-color: white;
    box-shadow: 0 0 5px rgba(0,0,0,0.5);
    border-radius: 10px;
    padding: 20px;
    font-size: 1.2em;
  }

  .list_form_header,.list_items,.add_items{
    width: 100%;
  }
  .list_form_header{
    /*padding: 10px;*/
  }
  .list_header,.add_input{
    width: 100%;
    border: none;
  }
  .list_header:focus,.add_input:focus{
    outline: none !important;
  }
  input:focus{

  }
</style>