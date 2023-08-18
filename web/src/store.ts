import { createStore } from 'ice';
import user from '@/models/user';
import app from '@/models/app';

export default createStore({
  user,
  app,
});
