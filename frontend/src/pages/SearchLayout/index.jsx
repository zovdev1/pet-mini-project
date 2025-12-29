import { Outlet } from 'react-router-dom'
import css from './index.module.css'
import { useState } from 'react'
import { SingUp } from '../../components/SingUp'
import { SingIn } from '../../components/SingIn'

export const SearchLayout = () => {
  const [searchValue, setSearchValue] = useState('')
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [SingInModalOpen, setSingInModalOpen] = useState(false)

  const toggleModal = () => {
    const nextState = !isModalOpen;
    setIsModalOpen(nextState);
    
    document.body.style.overflow = nextState ? 'hidden' : 'unset';
  };

  const handleSearchChange = (event) => {
    setSearchValue(event.target.value)
  }
  
  return (
    <div>
      <header className={css.container}>
        <a href="/">marketplace</a>

        <input
          className={css.serch}
          type="search"
          name="search"
          placeholder="Поиск..."
          value={searchValue}
          onChange={handleSearchChange}
        />

        <div className={css.login} onClick={toggleModal}>
          <p>войти</p>
        </div>

        <a href="/create">Сreate</a>
      </header>
      
      {isModalOpen && (
        <SingUp 
        onClose={() => {
          setIsModalOpen(false); 
          setSingInModalOpen(true); 
        }} 
        onSwitch={() => { 
          setIsModalOpen(false); 
          setSingInModalOpen(false); 
        }}/>
        )}

      {SingInModalOpen && (
        <SingIn 
        onClose={() => {
          setIsModalOpen(false); 
          setSingInModalOpen(true); 
        }} 
        onSwitch={() => { 
          setIsModalOpen(false); 
          setSingInModalOpen(false); 
        }}/>
      )}

      <main>
        <Outlet />
      </main>
    </div>
  )
}