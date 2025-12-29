import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { SearchLayout } from '../pages/SearchLayout'
import { HomePage } from '../pages/HomePage'
import { CartPage } from '../pages/CartPage'
import { CatalogPage } from '../pages/CatalogPage'

export const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<SearchLayout />}>
          <Route path="/" element={<HomePage />} />
          <Route path='/:id' element={<CartPage />}/>
          <Route path='/create' element={<CatalogPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}
