import { useState } from 'react';
import css from './index.module.css'
import z, { string } from 'zod'
import axios from 'axios'

const baseURL = 'http://127.0.0.1:80'

const signUpSchema = z.object({
  email: string(),
  password: string().min(3),
})

export const SingUp = ({ onClose, onSwitch }) => {
    const [loading, setLoading] = useState(false)
    const [errors, setErrors] = useState({})
    const [formData, setformData] = useState({
        email: '',
        password: '',
    })

    const handleRegisterSubmit = async (e) => {
        e.preventDefault();
        
        try {
            
            signUpSchema.parse(formData)

            setLoading(true)

            const response = await axios.post(`${baseURL}/api/v1/user/create`, formData)

            console.log(response);

        } finally {
          setLoading(false)
        }

        onSwitch();
        onClose();
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
            <h2>SingUp</h2>
            <form className={css.modalForm} onSubmit={handleRegisterSubmit}>
              <label>email</label>
              <input 
                    className={css.modalFormInput} 
                    onChange={handleChange} 
                    value={formData.email}
                    name='email'
                    type="email" 
                    placeholder="Email"
                    required
                />
              <label>password</label>
              <input 
                    className={css.modalFormInput} 
                    onChange={handleChange}
                    value={formData.password}
                    name='password' 
                    type="password" 
                    placeholder="password"
                    required
                />
              <button className={css.modalFormButton} onClick={onClose} disabled={loading} type="submit">Зарегистрироваться</button>
            </form>
            <button className={css.closeBtn} onClick={onSwitch}>×</button>
          </div>
        </div>
    )
}