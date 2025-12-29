import axios from 'axios'
import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { z } from 'zod'
import css from './index.module.css'

const baseURL = 'http://127.0.0.1:80'

const productFormSchema = z.object({
  id: z.uuid(),
  title: z.string(),
  description: z.string(),
  price: z.number(),
  quantity: z.number(),
  created_at: z.string(),
  updated_at: z.string(),
})

export const CartPage = () => {
  const { id } = useParams()
  const [data, setData] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(`${baseURL}/api/v1/product/${id}`)

        const validatedData = productFormSchema.parse(response.data)

        setData(validatedData)
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

    fetchData()
  }, [id])

  if (loading) {
    return <div>Загрузка...</div>
  }

  if (!data) {
    return <div>Нет данных</div>
  }

  return (
    <div className={css.container_main}>
      <Link to="/">
        / home
      </Link>
      <div className={css.grup_card}>
        <div className={css.car_img}>img</div>
        <div className={css.box_div}>
          <h1>{data.title}</h1>
          <p style={{margin: "10px 0 10px 0"}}>{data.description}</p>
          <p className={css.price_card}>{data.price.toLocaleString('ru-RU')} ₽ </p>
        </div>
      </div>
    </div>
  )
}
