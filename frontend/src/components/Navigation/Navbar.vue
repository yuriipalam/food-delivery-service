<script setup>
import NavLink from "./NavLink.vue";
import PrimaryButton from "../PrimaryButton.vue";
import {useRouter} from "vue-router";
import {useAuthStore} from "../../store";
import {onMounted, ref} from "vue";

const props = defineProps({
  isLight: Boolean,
})

const useAuth = useAuthStore()

const router = useRouter()

function getFill(isLight) {
  if (isLight) {
    return 'var(--milky)'
  }
  return 'var(--blackish)'
}

function hamburgerMenu() {
  const hamburger = document.querySelector('.hamburger')
  const ul = document.querySelector('ul')

  hamburger.addEventListener('click', () => {
    hamburger.classList.toggle('active')
    ul.classList.toggle('active')
  })

  document.querySelectorAll('a, #li-sign-in-button, #li-profile-button').forEach(link => {
    link.addEventListener('click', () => {
      hamburger.classList.remove('active')
      ul.classList.remove('active')
    })
  })
}

const currWidth = ref(0)

function isSmallScreen() {
  return currWidth.value <= 980
}

onMounted(() => {
  hamburgerMenu()

  currWidth.value = window.innerWidth
  window.addEventListener('resize', () => {
    currWidth.value = window.innerWidth
  })
})
</script>

<template>
  <nav>
    <div class="container">
      <NavLink :route-name="'Home'">
        <svg xmlns="http://www.w3.org/2000/svg" width="100" viewBox="0 0 294.079 222.398">
          <title>logo_food</title>
          <g>
            <path d="M190.333,187.922" transform="translate(-98.426 -49.965)" fill="#fff" stroke="#000" stroke-miterlimit="10"/>
            <path d="M120.386,178.166c4.122,3.061,8.645,5.93,13.159,8.624q6.8,4.056,13.92,7.594a232.359,232.359,0,0,0,29.275,12.254,207.938,207.938,0,0,0,62.354,11.145,196.314,196.314,0,0,0,31.726-2.232l7.843-1.409c1.316-.2,2.6-.56,3.89-.862l3.874-.938,3.876-.939c1.284-.341,2.544-.768,3.817-1.148l7.625-2.353A217.539,217.539,0,0,0,358,178a218.124,218.124,0,0,1-55.873,30.986l-7.629,2.506c-1.276.405-2.538.858-3.825,1.225l-3.887,1.016-3.888,1.016c-1.3.328-2.585.714-3.907.939l-7.887,1.563a198.1,198.1,0,0,1-32.028,2.831,210.845,210.845,0,0,1-63.415-10.168A236.119,236.119,0,0,1,145.694,198q-7.291-3.465-14.318-7.48a161.639,161.639,0,0,1-13.762-8.682,2.3,2.3,0,0,1,2.674-3.74l.024.017Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
            <g>
              <path d="M124.822,126.4q-2.441,0-3.008,1.238a15.036,15.036,0,0,0-.795,3.61,35.912,35.912,0,0,0-.372,4.882,1.909,1.909,0,0,0,.849,1.805,6.99,6.99,0,0,0,3.185.5,3.791,3.791,0,0,0,2.3-.566,3.1,3.1,0,0,0,1.026-1.91q.813-3.858,3.326-3.857,4.317,0,4.317,10.332,0,6.547-1.876,9.624a3.482,3.482,0,0,1-3.113,1.876q-2.088,0-2.407-2.8a12.055,12.055,0,0,0-.441-2.547,3.561,3.561,0,0,0-1.009-1.415,4.786,4.786,0,0,0-3.4-1.1,3.075,3.075,0,0,0-2.158.567,2.885,2.885,0,0,0-.5,2.158q0,.284-.035,3.291,0,6.121,1.274,7.466a2.048,2.048,0,0,0,.866.69,17.166,17.166,0,0,0,2.318.442,2.4,2.4,0,0,1,1.557.956,2.955,2.955,0,0,1,.637,1.875q0,2.726-3.184,3.5a58.742,58.742,0,0,1-12.916,1.1q-7.29,0-10.049-.849a3.96,3.96,0,0,1-2.017-1.416,3.656,3.656,0,0,1-.778-2.264,3.128,3.128,0,0,1,1.132-2.513,13.375,13.375,0,0,1,1.946-.884q1.981-.779,2.371-3.609.636-4.353.637-19.851a49.488,49.488,0,0,0-.46-8.669,3.768,3.768,0,0,0-2.23-2.76,7.709,7.709,0,0,1-2.317-1.362,2.365,2.365,0,0,1-.548-1.681,3.075,3.075,0,0,1,1.521-2.6,6.473,6.473,0,0,1,3.751-1.044q1.414,0,4.14.213,5.766.459,13.8.46,12.844,0,17.3-.708a8.538,8.538,0,0,1,.92-.071,3.126,3.126,0,0,1,2.37,1.415,13.873,13.873,0,0,1,3.114,8.776,5.525,5.525,0,0,1-1.026,3.5,3.254,3.254,0,0,1-2.689,1.309,3.016,3.016,0,0,1-1.876-.5,40.281,40.281,0,0,1-3.255-3.468,7.383,7.383,0,0,0-3.839-2.406A45.982,45.982,0,0,0,124.822,126.4Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
              <path d="M164.913,132.127a19.536,19.536,0,0,1,11.358,3.362,18,18,0,0,1,5.626,6.262,17.394,17.394,0,0,1,2.194,8.422q0,7.678-7.041,13.057a24.694,24.694,0,0,1-15.392,5.131q-8.457,0-13.694-4.813a15.742,15.742,0,0,1-5.095-12.172,17.124,17.124,0,0,1,6.227-13.658A23.451,23.451,0,0,1,164.913,132.127Zm-2.83,8.174a3.056,3.056,0,0,0-2.5,1.15,4.692,4.692,0,0,0-.938,3.061,29.016,29.016,0,0,0,.8,5.785,43.876,43.876,0,0,0,1.893,6.493q1.1,2.9,3.574,2.9,3.148,0,3.15-4a33.217,33.217,0,0,0-.8-6.758,27.867,27.867,0,0,0-1.928-6.157A3.748,3.748,0,0,0,162.083,140.3Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
              <path d="M208.541,132.127a19.537,19.537,0,0,1,11.359,3.362,18.012,18.012,0,0,1,5.627,6.262,17.416,17.416,0,0,1,2.194,8.422q0,7.678-7.041,13.057a24.7,24.7,0,0,1-15.393,5.131q-8.459,0-13.694-4.813a15.742,15.742,0,0,1-5.095-12.172,17.124,17.124,0,0,1,6.227-13.658A23.448,23.448,0,0,1,208.541,132.127Zm-2.83,8.174a3.056,3.056,0,0,0-2.494,1.15,4.692,4.692,0,0,0-.938,3.061,28.942,28.942,0,0,0,.8,5.785,43.876,43.876,0,0,0,1.893,6.493q1.1,2.9,3.573,2.9,3.15,0,3.15-4a33.21,33.21,0,0,0-.8-6.758,27.814,27.814,0,0,0-1.928-6.157A3.747,3.747,0,0,0,205.711,140.3Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
              <path d="M271.632,121.937v32.3a16.319,16.319,0,0,0,.229,3.539,3.068,3.068,0,0,0,1.292,1.521,6.557,6.557,0,0,1,1.716,1.557,2.9,2.9,0,0,1,.443,1.628q0,2.832-5.7,4.317a41.406,41.406,0,0,1-10.261,1.309,2.2,2.2,0,0,1-1.469-.336,2.847,2.847,0,0,1-.443-1.61q-.177-1.8-1.273-1.805a3.747,3.747,0,0,0-1.876,1.062,12.413,12.413,0,0,1-8.032,2.937,14.152,14.152,0,0,1-9.872-4.069,19.967,19.967,0,0,1-4.175-5.626,18.2,18.2,0,0,1-2.123-8.387,18.557,18.557,0,0,1,2.76-9.624,17.241,17.241,0,0,1,6.935-6.582,15.649,15.649,0,0,1,7.254-1.91A14.7,14.7,0,0,1,254.4,134a1.968,1.968,0,0,0,.85.284q1.133,0,1.132-3.043a2.9,2.9,0,0,0-.619-2,4.963,4.963,0,0,0-2.211-1.185q-2.231-.708-2.23-2.477,0-2.583,5.131-4.034,9.057-2.511,12.137-2.512Q271.631,119.035,271.632,121.937Zm-19.214,17.8a4.066,4.066,0,0,0-3.645,2.441,17.622,17.622,0,0,0-1.8,8.563q0,4.814,1.734,7.254a3.892,3.892,0,0,0,3.361,1.769q2.3,0,3.344-2.671t1.044-8.616q0-5.449-1.238-7.254A3.329,3.329,0,0,0,252.418,139.735Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
              <path d="M306.662,118.362a47.5,47.5,0,0,1,8.067.708,12.687,12.687,0,0,1,4.813,1.557,3.129,3.129,0,0,1,1.451,2.725,2.554,2.554,0,0,1-.531,1.716,7.962,7.962,0,0,1-2.3,1.5,3.97,3.97,0,0,0-2.353,2.494q-.3,1.432-.548,11.129-.108,9.377-.2,11.252a24.523,24.523,0,0,1-.442,3.644,14.716,14.716,0,0,1-4.141,8.1,17.839,17.839,0,0,1-6.545,3.786,25.612,25.612,0,0,1-8.422,1.381,31.112,31.112,0,0,1-7.872-1.027,24.628,24.628,0,0,1-6.848-2.83,11.15,11.15,0,0,1-5.236-9.413,8.113,8.113,0,0,1,2.228-5.856,7.435,7.435,0,0,1,5.591-2.318,8.367,8.367,0,0,1,5.591,1.947,6.2,6.2,0,0,1,2.3,4.883,12.061,12.061,0,0,1-.424,2.477,3.907,3.907,0,0,0-.072.778,2.322,2.322,0,0,0,.815,1.84,3.17,3.17,0,0,0,2.158.708,4.436,4.436,0,0,0,3.821-2.123,9.734,9.734,0,0,0,1.452-5.591l-.037-3.22-.07-8.669q-.318-7.89-.477-9.288a4.787,4.787,0,0,0-.8-2.283,3.17,3.17,0,0,0-.938-.919,14.248,14.248,0,0,0-2.211-.673,2.4,2.4,0,0,1-1.451-.991,2.862,2.862,0,0,1-.6-1.733,3.66,3.66,0,0,1,1.5-3.061,12.5,12.5,0,0,1,4.9-1.858A38.806,38.806,0,0,1,306.662,118.362Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
              <path d="M355.528,151.553H338.9q-1.1,0-1.1.955,0,2.3,2.494,4.335a8.324,8.324,0,0,0,5.361,2.034,12.584,12.584,0,0,0,3.768-.654,12.253,12.253,0,0,0,3.415-1.61,3.739,3.739,0,0,1,1.911-.849,2.409,2.409,0,0,1,1.8,1.132,3.756,3.756,0,0,1,.884,2.335q0,2.549-3.184,4.919a21.08,21.08,0,0,1-12.985,4.211,18.441,18.441,0,0,1-12.456-4.317,17.719,17.719,0,0,1-4.034-5.131,17.945,17.945,0,0,1-2.088-8.421,17.2,17.2,0,0,1,3.61-10.686,18.97,18.97,0,0,1,7.926-6.122,19.927,19.927,0,0,1,7.82-1.521,17.55,17.55,0,0,1,11.429,3.821,15.3,15.3,0,0,1,3.821,4.529,11.848,11.848,0,0,1,1.663,5.8,6.786,6.786,0,0,1-.955,3.768A2.86,2.86,0,0,1,355.528,151.553Zm-15.711-6.617h3.575q2.262,0,2.264-2.087a4.444,4.444,0,0,0-1.1-3.114,3.7,3.7,0,0,0-2.867-1.2,3.641,3.641,0,0,0-3.113,1.556,4.941,4.941,0,0,0-1.061,2.937,1.841,1.841,0,0,0,.512,1.487A2.863,2.863,0,0,0,339.817,144.936Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
              <path d="M381.819,141.433v11.96a5.492,5.492,0,0,0,1.044,3.468,3.239,3.239,0,0,0,2.671,1.345,8.094,8.094,0,0,0,3.645-1.274,2.043,2.043,0,0,1,1.061-.319,1.955,1.955,0,0,1,1.558,1.026,3.856,3.856,0,0,1,.707,2.229q0,2.2-2.406,4.211a18.134,18.134,0,0,1-12.172,4.282,12.793,12.793,0,0,1-9.094-3.2,11,11,0,0,1-3.433-8.439V140.372a.917.917,0,0,0-1.2-1.2h-3.644c-.685,0-1.12-.111-1.309-.335a2.792,2.792,0,0,1-.283-1.575V135.7a1.789,1.789,0,0,1,.813-1.557l16.206-10.191a3.157,3.157,0,0,1,1.487-.318h2.866a1.41,1.41,0,0,1,1.186.407,2.953,2.953,0,0,1,.3,1.61v4.777a1.977,1.977,0,0,0,.336,1.38,2.212,2.212,0,0,0,1.433.318h6.547a2.19,2.19,0,0,1,1.7.478,3.544,3.544,0,0,1,.389,2.034v2.088a4.076,4.076,0,0,1-.5,2.353,2.138,2.138,0,0,1-1.8.655h-6.4a1.379,1.379,0,0,0-1.7,1.7Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
            </g>
            <path d="M337.586,86.414a74.439,74.439,0,0,0-11.377-8.857c-4.1-2.644-8.392-5-12.753-7.237A176.546,176.546,0,0,0,286.07,59.555a171.916,171.916,0,0,0-58.33-7.183A165.9,165.9,0,0,0,198.5,56.534a159.325,159.325,0,0,0-28.038,9.446l-3.386,1.51-1.694.755c-.562.257-1.1.559-1.654.837l-6.593,3.4c-2.182,1.162-4.274,2.489-6.417,3.725l-3.2,1.888-3.089,2.058c-2.053,1.383-4.126,2.739-6.163,4.146l-5.936,4.469c-.984.752-1.992,1.474-2.956,2.252l-2.843,2.393-5.688,4.79,5.607-4.886,2.8-2.444c.951-.8,1.948-1.533,2.92-2.3L138.043,84c2.017-1.443,4.071-2.837,6.1-4.258l3.064-2.117,3.174-1.947c2.128-1.277,4.207-2.644,6.378-3.849l6.565-3.529c.55-.288,1.089-.6,1.649-.87l1.689-.787,3.379-1.576a160.6,160.6,0,0,1,28.088-9.994,167.746,167.746,0,0,1,29.487-4.7,174.429,174.429,0,0,1,59.29,6.293A179.737,179.737,0,0,1,315,67.156c8.983,4.449,17.828,9.3,25.41,16.43Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
          </g>
          <g>
            <path d="M138.046,217.291c-9.939,16.838-8.868,36.947,5.2,46.789C151.41,269.794,163.754,275.8,175,270c6.672-3.442,6-9,4-21-2.795-16.77.521-33.956,12.316-34.963,1.79-.153-21.246-11.421-37.316-8.037A24.957,24.957,0,0,0,138.046,217.291Z" transform="translate(-98.426 -49.965)" fill="#fe724c"/>
            <g>
              <path d="M373.452,112.841a10.262,10.262,0,0,1-7.2-1.738l-5.958-3.724-10.178,7.447-4.468,2.73s-1.241,1.986-1.986-1.241,9.93-15.887,9.93-15.887l-6.7-6.95-5.213-.5,4.468-2.731" transform="translate(-98.426 -49.965)" fill="#fe724c" stroke="#fe724c" stroke-miterlimit="10"/>
              <path d="M374.49,111.69a10.258,10.258,0,0,0-3.372-6.594l-5.013-4.922,4.862-11.636,1.61-4.983s1.641-1.671-1.671-1.641S357.78,95.282,357.78,95.282l-8.325-4.892-1.7-4.952-1.61,4.983" transform="translate(-98.426 -49.965)" fill="#fe724c" stroke="#fe724c" stroke-miterlimit="10"/>
            </g>
            <path d="M171,210l-14,6s-7,2-8,12c-.7,6.965,4,21,4,21l.854-16.226a18.464,18.464,0,0,1,1.314-6.795C157,222,158,221,159,220Z" transform="translate(-98.426 -49.965)" fill="#f3cf79" opacity="0.6"/>
          </g>
        </svg>

      </NavLink>
      <ul>
        <li>
          <NavLink :route-name="'Suppliers'" :is-light="props.isLight">Suppliers</NavLink>
        </li>
        <li>
          <NavLink :route-name="'Categories'" :is-light="props.isLight">Categories</NavLink>
        </li>
        <li v-if="isSmallScreen()">
          <PrimaryButton id="li-sign-in-button" v-if="useAuth.idRef === -1" @click="router.push({name: 'SignIn'})" :class="'header-button'">
            Sign
            in
          </PrimaryButton>
          <PrimaryButton id="li-profile-button" v-if="useAuth.idRef !== -1" @click="router.push({name: 'Profile'})" :class="'header-button'">
            Profile
          </PrimaryButton>
        </li>
      </ul>
      <div class="navbar-content">
        <NavLink :route-name="'Cart'" class="cart-link">
          <svg xmlns="http://www.w3.org/2000/svg" width="26" viewBox="0 0 29 33" :fill="getFill(props.isLight)">
            <path
                d="M28.509 4.2C28.3679 4.00623 28.1937 3.85081 27.9981 3.74422C27.8025 3.63763 27.59 3.58235 27.375 3.58209H6.6465L6.354 1.49731C6.29577 1.0791 6.11518 0.699114 5.84437 0.424943C5.57355 0.150771 5.23006 0.000195169 4.875 0H1.5C1.10218 0 0.720644 0.188699 0.43934 0.524585C0.158035 0.860471 0 1.31603 0 1.79104C0 2.26606 0.158035 2.72162 0.43934 3.0575C0.720644 3.39339 1.10218 3.58209 1.5 3.58209H3.6045L6.3945 23.5773L6.462 23.7994L6.543 24.0699L6.723 24.3904L6.8655 24.591L7.155 24.8239L7.323 24.9403C7.49745 25.0274 7.68443 25.073 7.8735 25.0746H24.375C24.7728 25.0746 25.1544 24.8859 25.4357 24.55C25.717 24.2142 25.875 23.7586 25.875 23.2836C25.875 22.8086 25.717 22.353 25.4357 22.0171C25.1544 21.6812 24.7728 21.4925 24.375 21.4925H9.1455L8.8965 19.7015H25.875C26.236 19.7016 26.5848 19.5462 26.8576 19.2639C27.1304 18.9817 27.3087 18.5914 27.36 18.1648L28.86 5.62746C28.8904 5.3734 28.8748 5.11451 28.8142 4.86826C28.7537 4.62201 28.6496 4.39413 28.509 4.2ZM25.6455 7.16418L25.218 10.7463H19.875V7.16418H25.6455ZM18.375 7.16418V10.7463H13.875V7.16418H18.375ZM18.375 12.5373V16.1194H13.875V12.5373H18.375ZM12.375 7.16418V10.7463H7.875L7.653 10.8L7.146 7.16418H12.375ZM7.896 12.5373H12.375V16.1194H8.3955L7.896 12.5373ZM19.875 16.1194V12.5373H25.002L24.5745 16.1194H19.875Z"
            />
            <path
                d="M10.125 32.2388C11.3676 32.2388 12.375 31.036 12.375 29.5522C12.375 28.0685 11.3676 26.8657 10.125 26.8657C8.88236 26.8657 7.875 28.0685 7.875 29.5522C7.875 31.036 8.88236 32.2388 10.125 32.2388Z"
            />
            <path
                d="M23.625 32.2388C24.8676 32.2388 25.875 31.036 25.875 29.5522C25.875 28.0685 24.8676 26.8657 23.625 26.8657C22.3824 26.8657 21.375 28.0685 21.375 29.5522C21.375 31.036 22.3824 32.2388 23.625 32.2388Z"
            />
          </svg>
        </NavLink>
        <PrimaryButton v-if="useAuth.idRef === -1 && !isSmallScreen()" @click="router.push({name: 'SignIn'})"
                       :class="'header-button'">Sign
          in
        </PrimaryButton>
        <PrimaryButton v-if="useAuth.idRef !== -1 && !isSmallScreen()" @click="router.push({name: 'Profile'})"
                       :class="'header-button'">
          Profile
        </PrimaryButton>
        <div class="hamburger">
          <span class="hamburger-bar"></span>
          <span class="hamburger-bar"></span>
          <span class="hamburger-bar"></span>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
nav {
  padding-top: 20px;
  padding-bottom: 20px;
}

svg, a, .hamburger {
  z-index: 10;
}

nav > .container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navbar-content {
  display: flex;
  align-items: center;
}

ul {
  margin: 0 0;
  padding-left: 0;
  list-style: none;
  display: flex;
}

ul li {
  margin-right: 100px;
}

ul li:last-child {
  margin-right: 0;
}

svg {
  vertical-align: middle;
}

.cart-link {
  margin-right: 50px;
}

.hamburger {
  display: none;
  cursor: pointer;
}

.hamburger-bar {
  display: block;
  width: 25px;
  height: 3px;
  margin: 5px auto;
  transition: all 0.3s ease-in-out;
  -webkit-transition: all 0.3s ease-in-out;
  background-color: v-bind('getFill(props.isLight)');
}

@media screen and (max-width: 980px) {
  .hamburger {
    display: block;
  }

  .hamburger.active .hamburger-bar:nth-child(1) {
    transform: translateY(8px) rotate(45deg);
  }

  .hamburger.active .hamburger-bar:nth-child(2) {
    opacity: 0;
  }

  .hamburger.active .hamburger-bar:nth-child(3) {
    transform: translateY(-8px) rotate(-45deg);
  }

  ul {
    position: absolute;
    z-index: 1;
    top: -100%;
    left: 0;
    gap: 0;
    flex-direction: column;
    background-color: v-bind('getFill(!props.isLight)');
    width: 100%;
    text-align: right;
    align-items: flex-end;
    transition: 0.7s;
    padding-right: 25px;
  }

  ul li, ul li:last-child {
    margin: 16px 40px;
  }

  ul.active {
    padding-top: 75px;
    top: 0;
  }
}

@media screen and (max-width: 780px) {
  ul li, ul li:last-child {
    margin: 16px 28px;
  }
}

@media screen and (max-width: 480px) {
  ul li, ul li:last-child {
    margin: 16px 20px;
  }
}
</style>
