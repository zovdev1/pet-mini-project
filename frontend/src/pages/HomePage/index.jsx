import axios from 'axios'
import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { z } from 'zod'
import css from './index.module.css'

const baseURL = 'http://127.0.0.1:80'

const productFormSchema = z.array(
  z.object({
    id: z.uuid(),
    title: z.string(),
    description: z.string(),
    price: z.number(),
    quantity: z.number(),
    created_at: z.string(),
    updated_at: z.string(),
  })
)
export const HomePage = () => {
  const [data, setData] = useState([])
  const [loading, setLoading] = useState(true)
  const [limit, setLimit] = useState(15)
  const [isFetching, setIsFetching] = useState(false)

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true)

        const response = await axios.get(`${baseURL}/api/v1/product/?limit=${limit}&offset=0`)
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
        setIsFetching(false)
      }
    }

    fetchData()
  }, [limit])

  useEffect(() => {
    const onScroll = () => {
      const bottom = window.innerHeight + window.scrollY >= document.body.offsetHeight - 100

      if (bottom && !isFetching) {
        setIsFetching(true)
        setLimit((prev) => prev + 10)
      }
    }

    window.addEventListener('scroll', onScroll)
    return () => window.removeEventListener('scroll', onScroll)
  }, [isFetching])

  if (loading && data.length === 0) {
    return <div>Загрузка...</div>
  }

  return (
    <div className={css.container}>
      {data.map((item) => (
        <div key={item.id} className={css.card}>
          <Link to={`/${item.id}`} className={css.cardLink}>
            <div className={css.product_amg}>img</div>
            <div className={css.grub}>
              <h3 className={css.title}>{item.title}</h3>
              <p className={css.description}>{item.description}</p>
              <p className={css.price}>{item.price.toLocaleString('ru-RU')} RUB</p>
            </div>
          </Link>
        </div>
      ))}

      {loading && <div style={{ textAlign: 'center', padding: 20 }}>Загрузка...</div>}
    </div>
  )
}
