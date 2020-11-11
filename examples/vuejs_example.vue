<template>
    <div id="contact">
  
      <div class="form" >

  
          <form @submit.prevent="submit">
  
              <div class="form-overlay" v-if="loading">
                  <font-awesome-icon icon="circle-notch" spin />
              </div><!-- form-overlay -->
  
              <div>
                  <div class="status" v-if="status !== ''">
                      <span v-if="status === 'success'">Your email has successfully been sent.</span>
                      <span v-if="status === 'error'">Something went wrong while sending your email.</span>
                  </div>
              </div>
  
              <div>
                  <input type="text" name="name" v-model="form_data.name" placeholder="Name or Company" required />
              </div>
  
              <div>
                  <input type="email" name="email" v-model="form_data.email" placeholder="Email" required />
              </div>
  
              <div>
                  <input type="text" name="subject" v-model="form_data.subject" placeholder="Subject" required />
              </div>
  
  
              <div>
                  <textarea type="text" name="content" v-model="form_data.content" placeholder="Message" rows="6" required />
              </div>
  
  
              <div>
                  <button type="submit">Send</button>
              </div>
          
        </form>
  
      </div>
  
    </div><!-- contact -->
</template>
  
<script>

import config from '../config'


export default {
    name: 'Contact',
    components: {
    },
    data: function() {
        return {
            contact_text: "",
            form_data: {
                name: "",
                email: "",
                subject: "",
                content: ""
            },
            status: "",
            loading: false
        }
    },
    created() {
    },
    mounted() {
    },
    methods: {
        clearForm: function() {
            this.form_data.name    = "";
            this.form_data.email   = "";
            this.form_data.subject = "";
            this.form_data.content = "";
        },
        submit: function(e) {
            e.preventDefault();

            var self = this
            self.loading = true
            
            this.axios.post(config.MAILBEAR_URL + `/api/v1/form/10810dce-1074-4988-a8f5-4c538a749a95`, this.form_data)
            .then(response => {
                self.status = "success"
                self.clearForm()

                return response
            })
            .catch(error => {
                self.status = "error"
                console.log(error)
            })
            .then(function () {
                // always executed
                self.loading = false
            })
        }
    }
}
</script>

<style lang="scss">


.form {
    max-width: 700px;
    margin: auto;

    .contact-text {
    padding: 1em;

    a {
        color: inherit
    }
    }

}


form {

    padding: 1em;

    position: relative;

    .form-overlay {
        position: absolute;
        background-color: rgba(255,255,255, 0.8);
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        font-size: 1.6em;
        display: flex;
        align-items: center;
        justify-content: center;

        
    }

    div {

        padding: 0.5em 0;

        input,
        textarea {
            border-width: 1px;
            border-radius: 0px;
            border-color: rgb(44, 62, 80);

            font-family: inherit;
            font-size: 1.2em;

            width: 100%;
            margin: 0;
            padding: 0.5em;
            box-sizing: border-box;
        }

        button {
            background-color: rgb(44, 62, 80);
            color: #fff;
            margin: 0;
            padding: 0.5em;
            border-width: 0;

            font-family: inherit;
            font-size: 1.2em;
        }

        .status {
            background-color: rgb(44, 62, 80);
            color: #fff;
            margin: 0;
            padding: 0.5em;
        }

    }

    

}


</style>
