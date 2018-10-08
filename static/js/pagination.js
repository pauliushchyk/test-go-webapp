const createLink = (page) => {
  const params = window.location.search
    .replace('?', '')
    .split('&')
    .filter(param => !param.includes('page'))
    .join('&');

  const a = document.createElement('a');

  let query = `/customers?page=${page}`;

  if (params.length > 0) {
    query += `&${params}`;
  }

  a.setAttribute('href', query);
  a.text = page;

  return a;
}

const createSpan = () => {
  const span = document.createElement('span');

  span.innerText = '...';

  return span;
}

(() => {
  const element = document.getElementById('pagination');

  const pageParam = window.location.search
    .replace('?', '')
    .split('&')
    .find(param => param.includes('page'));
  const page = pageParam ? pageParam.replace('page=', '') : 0;

  const last = Math.ceil(element.dataset.count / 10);
  const delta = 2;
  const left = page - delta;
  const right = page + delta + 1;
  const range = [];
  let l;

  for (let i = 1; i <= last; i++) {
    if (i === 1 || i === last || i >= left && i < right) {
      range.push(i);
    }
  }

  for (let i of range) {
    if (l) {
      if (i - l === 2) {
        element.appendChild(createLink(l + 1));
      } else if (i - l !== 1) {
        element.appendChild(createSpan());
      }
    }

    l = i;

    element.appendChild(createLink(i));
  }
})();

(() => {
  const params = window.location.search
    .replace('?', '')
    .split('&');

  const currentParams = params.filter(param => !param.includes('by') && !param.includes('order')).join('&');

  const byParam = params.find(param => param.includes('by'));
  const by = byParam && byParam.replace('by=', '');

  const orderParam = params.find(param => param.includes('order'));

  const links = document.querySelectorAll('.customers tr th a[data-field]');

  links.forEach(link => {
    let query = `/customers?by=${link.dataset.field}`;

    if (currentParams.length > 0) {
      query += `&${currentParams}`;
    }

    if (link.dataset.field === by) {
      link.classList.add('active');

      if (!orderParam) {
        query += '&order=desc';
      }
    }

    link.setAttribute('href', query);
  })
})();