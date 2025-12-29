import { useState } from 'react';
import css from './index.module.css'
import axios from 'axios'
import z from 'zod'

const baseURL = 'http://127.0.0.1:80'

const LoginForm = z.object({
  email: z.string(),
  password: z.string(),
})


export const SingIn = ({ onClose, onSwitch }) => {
    const [loading, setLoading] = useState(false)
    const [errors, setErrors] = useState({})
    const [formData, setformData] = useState({
        email: '',
        password: '',
    })
    

    const handleRegisterSubmit = async (e) => {
        e.preventDefault();

        try {
          LoginForm.parse(formData)

          console.log(formData);

          setLoading(true)
          
          const response = await axios.post(`${baseURL}/api/v1/user/logIn`, formData)

          console.log(response);
          const token = response.data.token
          
          localStorage.setItem('authToken', token)
          
          axios.defaults.headers.common['Authorization'] = `Bearer ${token}`

        } finally {
          setLoading(false)
        }
    };

    const handleChange = (e) => {
        const { name, value } = e.target

        setformData((prev) => ({
            ...prev,
            [name]: value,
        }))

        if (errors[name]) {
            setErrors((prev) => ({
                ...prev,
                [name]: '',
            }))
        }
    }

    return (
        <div className={css.overlay}>
          <div className={css.modalContent} onClick={(e) => e.stopPropagation()}>
            <h2>SingIn</h2>
            <form className={css.modalForm} onSubmit={handleRegisterSubmit}>
              <label>email</label>
              <input 
                className={css.modalFormInput}
                onChange={handleChange}
                name='email'
                type="email" 
                placeholder="Email"
                value={formData.email}
              />
              <label>password</label>
              <input 
                className={css.modalFormInput}
                onChange={handleChange}
                name='password'
                type="password" 
                placeholder="password" 
                value={formData.password}
              />
              <button className={css.modalFormButton}  onClick={onClose} disabled={loading} type="submit">Зарегистрироваться</button>
            </form>
            <button className={css.closeBtn} onClick={onSwitch}>×</button>
          </div>
          singin
        </div>
    )
}