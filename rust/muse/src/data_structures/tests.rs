use super::*;

pub(crate) fn list_derive<'a, T>() -> impl Fn() + 'a
where
    T: Default + List<i32>,
{
    move || {
        let size: usize = 8192;

        let mut list = T::default();

        (1..=size).for_each(|i| {
            list.append(i as i32);
        });

        assert_eq!(list.size(), size);

        (0..size).for_each(|i| {
            assert_eq!(*list.get(i), i as i32 + 1);
        });

        (0..size).for_each(|i| {
            list.set(i, (size - i) as i32);
        });

        (0..size).for_each(|i| {
            assert_eq!(*list.get(i), (size - i) as i32);
        });

        let mut i = 0;
        let mut j = size - 1;
        while i < j {
            let x = *list.get(i);
            let y = *list.get(j);

            list.remove(i);
            list.insert(i, y);
            list.remove(j);
            list.insert(j, x);
            i += 1;
            j -= 1;
        }

        (0..size).for_each(|i| {
            assert_eq!(*list.get(i), i as i32 + 1);
        });

        (1..=size).rev().for_each(|i| {
            assert_eq!(*list.back(), i as i32);
            list.eject();
        });

        assert!(list.is_empty());
    }
}

pub(crate) fn list_derive_resizable<'a, T>() -> impl Fn() + 'a
where
    T: Default + List<i32> + Resizable,
{
    move || {
        let size: usize = 8192;

        let mut list = T::default();

        (1..=size).for_each(|i| {
            list.append(i as i32);
        });

        list.shrink();

        assert_eq!(list.capacity(), size);

        (1..=size).rev().for_each(|_| {
            list.eject();
        });

        list.shrink();

        assert_eq!(list.capacity(), 0);
    }
}
