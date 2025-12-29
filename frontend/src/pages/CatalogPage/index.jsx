import css from './index.module.css'
import axios from 'axios'
import { useState } from 'react'
import z from 'zod'

const baseURL = 'http://127.0.0.1:80'

const formDtataParse = z.object({
  title: z.string().min(3),
  description: z.string().min(3),
  price: z.number().min(0),
  quantity: z.number().min(0),
})

export const CatalogPage = () => {
    const [loading, setLoading] = useState(false)
    const [errors, setErrors] = useState({})
    const [formData, setformData] = useState({
        title: '',
        description: '',
        price: '',
        quantity: '',
    })

    const handleRegisterSubmit = async (e) => {
        e.preventDefault();

        try {

            const dataToValidate = {
                title: formData.title,
                description: formData.description,
                price: Number(formData.price) || 0,
                quantity: Number(formData.quantity) || 0,
            }

            formDtataParse.parse(dataToValidate)

            const token = localStorage.getItem('authToken')

            setLoading(true)

            const response = await axios.post(`${baseURL}/api/v1/product/`, dataToValidate, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })

            console.log(response);

        } catch (error) {
        if (error instanceof z.ZodError) {
          console.error('Ошибка валидации данных:', error.errors)
        } else {
          console.error('Ошибка загрузки:', error)
        }
        } finally {
            setLoading(false)
        }
    }

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
        <div className={css.container}>
            <form className={css.containerForm} onSubmit={handleRegisterSubmit}>
                <div>
                    <label> title </label>
                    <input
                        className={css.containerInput}
                        onChange={handleChange}
                        value={formData.title}
                        type="text"
                        name="title"
                        placeholder="title"
                    />
                </div>
                <div>
                    <label> description </label>
                    <input
                        className={css.containerInput}
                        onChange={handleChange}
                        value={formData.description}
                        type="text" 
                        name="description"
                        placeholder="description"
                    />
                </div>
                <div>
                    <label> price </label>
                    <input
                        className={css.containerInput}
                        onChange={handleChange}
                        value={formData.price}
                        type="text"
                        name="price"
                        placeholder="price"
                    />
                </div>
                <div>
                    <label> quantity </label>
                    <input
                        className={css.containerInput}
                        onChange={handleChange}
                        value={formData.quantity}
                        type="text"
                        name="quantity"
                        placeholder="quantity"
                    />
                </div>
                <button className={css.containerButton} disabled={loading} type="submit">but</button>
            </form>
        </div>
    )
}