import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Menu } from '../components/Menu';
import { Home } from './Home';
import { User } from './User';
import { Create } from './User/Create';

function Index() {
  return (
    <BrowserRouter>
      <Menu />
      <Routes>
        <Route path="/" element={<Home />}></Route>
        <Route path="/user" element={<User />}></Route>
        <Route path="/user/create" element={<Create />}></Route>
        <Route path="*" element={<Home />}></Route>
      </Routes>
    </BrowserRouter>
  );
}

export default Index;
